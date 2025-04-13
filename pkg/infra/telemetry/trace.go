package telemetry

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	otelTrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func (t *HttpTelemetry) Start(ctx context.Context, spanName string, opts ...otelTrace.SpanStartOption) (context.Context, otelTrace.Span) {
	return t.Tracer.Start(ctx, spanName, opts...)
}

func (t *HttpTelemetry) StartRoot(ctx context.Context, request *http.Request, spanName string) (context.Context, otelTrace.Span) {
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

func (t *HttpTelemetry) ShutdownTracer(ctx context.Context) error {
	return t.ShutdownTracerProvider(ctx)
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

func GetTraceID(ctx context.Context) string {
	spanCtx := otelTrace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		traceID := spanCtx.TraceID()
		return traceID.String()
	}

	var traceID [16]byte = [16]byte{
		0x00, 0x00, 0x00, 0x00,
	}
	spanCtx.WithTraceID(traceID)

	return ""
}
