package tracer

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/trace"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	otelTrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serviceName = semconv.ServiceNameKey.String("medzoner-service")

// Initialize a gRPC connection to be used by both the tracer and meter
// providers.
func initConn(host string) (*grpc.ClientConn, error) {
	// It connects the OpenTelemetry Collector through local gRPC connection.
	// You may replace `localhost:4317` with your endpoint.
	conn, err := grpc.NewClient(host,
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return conn, err
}

// Initializes an OTLP exporter, and configures the corresponding trace provider.
func initTracerProvider(ctx context.Context, res *resource.Resource, conn *grpc.ClientConn) (func(context.Context) error, error) {
	// Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	// Set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown will flush any remaining spans and shut down the exporter.
	return tracerProvider.Shutdown, nil
}

// Initializes an OTLP exporter, and configures the corresponding meter provider.
func initMeterProvider(ctx context.Context, res *resource.Resource, conn *grpc.ClientConn) (func(context.Context) error, error) {
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create metrics exporter: %w", err)
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(meterProvider)

	return meterProvider.Shutdown, nil
}

func initOtel(host string) (otelTrace.Tracer, metric.Meter, func(context.Context) error, func(context.Context) error) {
	log.Printf("Otel: Waiting for connection...")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	conn, err := initConn(host)
	if err != nil {
		log.Fatal(err)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			serviceName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	shutdownTracerProvider, err := initTracerProvider(ctx, res, conn)
	if err != nil {
		log.Fatal(err)
	}

	shutdownMeterProvider, err := initMeterProvider(ctx, res, conn)
	if err != nil {
		log.Fatal(err)
	}

	name := "go.opentelemetry.io/otel/example/otel-collector"
	tracer := otel.Tracer(name)
	meter := otel.Meter(name)

	// Attributes represent additional key-value descriptors that can be bound
	// to a metric observer or recorder.
	//commonAttrs := []attribute.KeyValue{
	//	attribute.String("attrA", "chocolate"),
	//	attribute.String("attrB", "raspberry"),
	//	attribute.String("attrC", "vanilla"),
	//}
	//
	//runCount, err := meter.Int64Counter("run", metric.WithDescription("The number of times the iteration ran"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Work begins
	//ctx, span := tracer.Start(
	//	ctx,
	//	"CollectorExporter-Example",
	//	otelTrace.WithAttributes(commonAttrs...))
	//defer span.End()
	//for i := 0; i < 10; i++ {
	//	_, iSpan := tracer.Start(ctx, fmt.Sprintf("Sample-%d", i))
	//	runCount.Add(ctx, 1, metric.WithAttributes(commonAttrs...))
	//	log.Printf("Doing really hard work (%d / 10)\n", i+1)
	//
	//	<-time.After(time.Second)
	//	iSpan.End()
	//}

	log.Printf("Done!")

	return tracer, meter, shutdownTracerProvider, shutdownMeterProvider
}

//go:generate mockgen -destination=../../../test/mocks/pkg/infra/tracer/http_tracer.go -package=tracerMock -source=./http_tracer.go Tracer
type Tracer interface {
	WriteLog(ctx context.Context, message string)
	Start(ctx context.Context, spanName string, opts ...otelTrace.SpanStartOption) (context.Context, otelTrace.Span)
	Int64Counter(name string, options ...metric.Int64CounterOption) (metric.Int64Counter, error)

	ShutdownTracer(ctx context.Context) error
	ShutdownMeter(ctx context.Context) error
}

type HttpTracer struct {
	Tracer                 otelTrace.Tracer
	Meter                  metric.Meter
	ShutdownTracerProvider func(context.Context) error
	ShutdownMeterProvider  func(context.Context) error
}

func (t HttpTracer) ShutdownTracer(ctx context.Context) error {
	return t.ShutdownTracerProvider(ctx)
}

func (t HttpTracer) ShutdownMeter(ctx context.Context) error {
	return t.ShutdownMeterProvider(ctx)
}

func (t HttpTracer) Start(ctx context.Context, spanName string, opts ...otelTrace.SpanStartOption) (context.Context, otelTrace.Span) {
	return t.Tracer.Start(
		ctx,
		spanName,
		opts...)
}

func (t HttpTracer) Int64Counter(name string, options ...metric.Int64CounterOption) (metric.Int64Counter, error) {
	return t.Meter.Int64Counter(name, options...)
}

func NewHttpTracer(config config.IConfig) (*HttpTracer, error) {
	f, err := os.Create(config.GetTraceFile())
	if err != nil {
		return nil, fmt.Errorf("failed to create trace output file: %v", err)
	}

	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("failed to close trace file: %v", err)
	}

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	tracer, meter, shutdownTracerProvider, shutdownMeterProvider := initOtel(config.GetOtelHost())
	return &HttpTracer{
		Tracer:                 tracer,
		Meter:                  meter,
		ShutdownTracerProvider: shutdownTracerProvider,
		ShutdownMeterProvider:  shutdownMeterProvider,
	}, nil
}

func prepWork() {
	fmt.Printf("this function will be traced\n")
}

func remainingWork() {
	fmt.Printf("this function will be traced2\n")
}

func (t HttpTracer) WriteLog(ctx context.Context, message string) {
	ctx, task := trace.NewTask(ctx, "awesomeTask")
	trace.Log(ctx, "orderID", message)
	trace.WithRegion(ctx, message, prepWork)
	// preparation of the task
	go func() { // continue processing the task in a separate goroutine.
		defer task.End()
		trace.WithRegion(ctx, message, remainingWork)
	}()
}
