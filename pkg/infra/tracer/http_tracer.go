package tracer

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	otelLog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	otelTrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	tracerName  = "medzoner/otel-collector"
	serviceName = "medzoner-service"
)

type Tracer interface {
	StartRoot(ctx context.Context, request *http.Request, spanName string) (context.Context, otelTrace.Span)
	Start(ctx context.Context, spanName string, opts ...otelTrace.SpanStartOption) (context.Context, otelTrace.Span)
	Error(span otelTrace.Span, err error) error
	ShutdownTracer(ctx context.Context) error
	ShutdownMeter(ctx context.Context) error
	ShutdownLogger(ctx context.Context) error
}

type HttpTracer struct {
	Tracer                 otelTrace.Tracer
	Meter                  metric.Meter
	ShutdownTracerProvider func(context.Context) error
	ShutdownMeterProvider  func(context.Context) error
	ShutdownLoggerProvider func(context.Context) error
	Logger                 *slog.Logger
}

func NewHttpTracer(cfg config.Config) (*HttpTracer, error) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	conn, err := initConn(cfg.OtelHost)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}

	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)))
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	shutdownTracerProvider, err := initTracerProvider(ctx, res, conn)
	if err != nil {
		return nil, err
	}

	shutdownMeterProvider, err := initMeterProvider(ctx, res, conn)
	if err != nil {
		return nil, err
	}

	shutdownLoggerProvider, err := initLogger(ctx, conn)
	if err != nil {
		return nil, err
	}

	return &HttpTracer{
		Tracer:                 otel.Tracer(tracerName),
		Meter:                  otel.Meter(tracerName),
		Logger:                 otelslog.NewLogger(serviceName),
		ShutdownTracerProvider: shutdownTracerProvider,
		ShutdownMeterProvider:  shutdownMeterProvider,
		ShutdownLoggerProvider: shutdownLoggerProvider,
	}, nil
}

func (t *HttpTracer) Start(ctx context.Context, spanName string, opts ...otelTrace.SpanStartOption) (context.Context, otelTrace.Span) {
	ctx, span := t.Tracer.Start(ctx, spanName, opts...)
	span.SetAttributes(attribute.String("correlation_id", middleware.GetCorrelationID(ctx)))
	return ctx, span
}

func (t *HttpTracer) StartRoot(ctx context.Context, request *http.Request, spanName string) (context.Context, otelTrace.Span) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	return t.Start(ctx, spanName,
		otelTrace.WithSpanKind(otelTrace.SpanKindServer),
		otelTrace.WithNewRoot(),
		otelTrace.WithAttributes(
			attribute.String("host", request.Host),
			attribute.String("path", request.URL.Path),
			attribute.String("method", request.Method),
		),
	)
}

func (t *HttpTracer) Error(span otelTrace.Span, err error) error {
	span.RecordError(err)
	t.Logger.Error(err.Error())
	return fmt.Errorf("error during event handling: %w", err)
}

func (t *HttpTracer) ShutdownLogger(ctx context.Context) error {
	return t.ShutdownLoggerProvider(ctx)
}

func (t *HttpTracer) ShutdownTracer(ctx context.Context) error {
	return t.ShutdownTracerProvider(ctx)
}

func (t *HttpTracer) ShutdownMeter(ctx context.Context) error {
	return t.ShutdownMeterProvider(ctx)
}

func initConn(host string) (*grpc.ClientConn, error) {
	return grpc.NewClient(host,
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}

func initTracerProvider(ctx context.Context, res *resource.Resource, conn *grpc.ClientConn) (func(context.Context) error, error) {
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("trace exporter creation failed: %w", err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(traceExporter)),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider.Shutdown, nil
}

func initMeterProvider(ctx context.Context, res *resource.Resource, conn *grpc.ClientConn) (func(context.Context) error, error) {
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("metrics exporter creation failed: %w", err)
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(meterProvider)

	return meterProvider.Shutdown, nil
}

func initLogger(ctx context.Context, conn *grpc.ClientConn) (func(context.Context) error, error) {
	logExporter, err := otlploggrpc.New(ctx, otlploggrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("logs exporter creation failed: %w", err)
	}

	loggerProvider := otelLog.NewLoggerProvider(otelLog.WithProcessor(otelLog.NewBatchProcessor(logExporter)))
	global.SetLoggerProvider(loggerProvider)

	return loggerProvider.Shutdown, nil
}
