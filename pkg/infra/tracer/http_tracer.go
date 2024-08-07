package tracer

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"log"
	"os"
	"runtime/trace"
)

type Tracer interface {
	WriteLog(ctx context.Context, message string)
}

type HttpTracer struct{}

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
	return &HttpTracer{}, nil
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
