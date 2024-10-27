package handler_test

import (
	"context"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"
	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/trace/noop"
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
	t.Run("Unit: test NotFoundHandler success", func(t *testing.T) {
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		notFoundHandler := handler.NotFoundHandler{
			Template: &TemplaterTest{},
			Tracer:   httpTracerMock,
		}
		request := httptest.NewRequest("GET", "/not-found", nil)
		notFoundHandler.Handle(httptest.NewRecorder(), request)
	})
	t.Run("Unit: test NotFoundHandler failed", func(t *testing.T) {
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		notFoundHandler := handler.NotFoundHandler{
			Template: &TemplaterTestFailed{},
			Tracer:   httpTracerMock,
		}
		request := httptest.NewRequest("GET", "/not-found", nil)

		notFoundHandler.Handle(httptest.NewRecorder(), request)
	})
}
