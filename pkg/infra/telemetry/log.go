package telemetry

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	otelTrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func initLogger(ctx context.Context, conn *grpc.ClientConn) (func(context.Context) error, error) {
	logExporter, err := newExporter(ctx, conn)
	if err != nil {
		return nil, fmt.Errorf("log exporter creation failed: %w", err)
	}

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
		log.WithResource(newResource()),
	)
	global.SetLoggerProvider(loggerProvider)

	return loggerProvider.Shutdown, nil
}

func newExporter(ctx context.Context, conn *grpc.ClientConn) (*otlploggrpc.Exporter, error) {
	exporter, err := otlploggrpc.New(ctx, otlploggrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("log exporter creation failed: %w", err)
	}
	return exporter, nil
}

type (
	LogLevel  string
	LogOutput string
)

const (
	TraceLevel   LogLevel = "trace"
	DebugLevel   LogLevel = "debug"
	InfoLevel    LogLevel = "info"
	WarningLevel LogLevel = "warning"
	ErrorLevel   LogLevel = "error"
	FatalLevel   LogLevel = "fatal"
	PanicLevel   LogLevel = "panic"
)

type Level LogLevel

func (l Level) AsZerolog() zerolog.Level {
	switch LogLevel(l) {
	case TraceLevel:
		return zerolog.TraceLevel
	case DebugLevel:
		return zerolog.DebugLevel
	case InfoLevel:
		return zerolog.InfoLevel
	case WarningLevel:
		return zerolog.WarnLevel
	case ErrorLevel:
		return zerolog.ErrorLevel
	case FatalLevel:
		return zerolog.FatalLevel
	case PanicLevel:
		return zerolog.PanicLevel
	}

	return zerolog.WarnLevel
}

// NewLoggerAdapter NewLoggerAdapter
func NewLoggerAdapter() (*zerolog.Logger, error) {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano
	zerolog.SetGlobalLevel(Level("debug").AsZerolog())

	hookedLogger := zerolog.New(zerolog.ConsoleWriter{
		Out: os.Stderr,
	}).With().CallerWithSkipFrameCount(3).Logger()

	return &hookedLogger, nil
}

func newResource() *resource.Resource {
	host, _ := os.Hostname()
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
		semconv.ServiceVersion("0.0.1"),
		semconv.HostName(host),
	)
}

func (t *HttpTelemetry) Log(ctx context.Context, msg string) {
	t.logger.Info(msg)

	t.zerolog.Info().
		Timestamp().
		Ctx(ctx).
		Msg(msg)
}

func (t *HttpTelemetry) Error(ctx context.Context, msg string, args ...any) {
	args = append([]any{"service.name", serviceName}, args...)
	args = append([]any{"service_name", serviceName}, args...)
	slog.ErrorContext(ctx, msg, args...)
	t.zerolog.Error().Timestamp().Ctx(ctx).Err(errors.New(msg)).Msg(msg)
}

func (t *HttpTelemetry) ErrorSpan(span otelTrace.Span, err error) error {
	span.RecordError(err)
	return err
}

func (t *HttpTelemetry) ShutdownLogger(ctx context.Context) error {
	return t.ShutdownLoggerProvider(ctx)
}
