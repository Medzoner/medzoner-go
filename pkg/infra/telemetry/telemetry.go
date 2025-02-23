package telemetry

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	otelTrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	tracerName  = "medzoner/otel-collector"
	serviceName = "medzoner-service"
)

//go:generate mockgen -destination=../../../test/mocks/telmeter.go -package=tracerMock -source=./telemetry.go
type Telemeter interface {
	StartRoot(ctx context.Context, request *http.Request, spanName string) (context.Context, otelTrace.Span)
	Start(ctx context.Context, spanName string, opts ...otelTrace.SpanStartOption) (context.Context, otelTrace.Span)
	ErrorSpan(span otelTrace.Span, err error) error
	ShutdownTracer(ctx context.Context) error
	ShutdownMeter(ctx context.Context) error
	ShutdownLogger(ctx context.Context) error
	Log(ctx context.Context, msg string)
	Error(ctx context.Context, msg string, args ...any)
}

type HttpTelemetry struct {
	Tracer                 otelTrace.Tracer
	Meter                  metric.Meter
	ShutdownTracerProvider func(context.Context) error
	ShutdownMeterProvider  func(context.Context) error
	ShutdownLoggerProvider func(context.Context) error
	logger                 *slog.Logger
	zerolog                *zerolog.Logger
}

func NewHttpTelemetry(cfg config.Config) (*HttpTelemetry, error) {
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

	zl, err := NewLoggerAdapter()
	if err != nil {
		return nil, fmt.Errorf("error creating logger: %w", err)
	}

	osl := otelslog.NewLogger(tracerName)
	slog.SetDefault(osl)

	return &HttpTelemetry{
		Tracer:                 otel.Tracer(tracerName),
		Meter:                  otel.Meter(tracerName),
		logger:                 osl,
		ShutdownTracerProvider: shutdownTracerProvider,
		ShutdownMeterProvider:  shutdownMeterProvider,
		ShutdownLoggerProvider: shutdownLoggerProvider,
		zerolog:                zl,
	}, nil
}

func initConn(host string) (*grpc.ClientConn, error) {
	cli, err := grpc.NewClient(host,
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	return cli, nil
}
