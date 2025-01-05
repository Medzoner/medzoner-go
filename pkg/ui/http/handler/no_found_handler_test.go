package handler_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/telemetry"

	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/trace/noop"
)

func TestNotFoundHandler(t *testing.T) {
	httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
	httpTelemetryMock.EXPECT().StartRoot(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	t.Run("Unit: test NotFoundHandler success", func(t *testing.T) {
		_ = t
		notFoundHandler := handler.NotFoundHandler{
			Template: &TemplaterTest{},
			Tracer:   httpTelemetryMock,
		}
		request := httptest.NewRequest("GET", "/not-found", nil)
		notFoundHandler.Handle(httptest.NewRecorder(), request)
	})
	t.Run("Unit: test NotFoundHandler failed", func(t *testing.T) {
		_ = t
		notFoundHandler := handler.NotFoundHandler{
			Template: &TemplaterTestFailed{},
			Tracer:   httpTelemetryMock,
		}
		request := httptest.NewRequest("GET", "/not-found", nil)

		notFoundHandler.Handle(httptest.NewRecorder(), request)
	})
}

type TemplaterTestFailed struct {
	RootPath string
}

func (t *TemplaterTestFailed) Render(name string, view interface{}, response http.ResponseWriter) (interface{}, error) {
	_ = name
	_ = view
	_ = response
	_ = t
	return nil, errors.New("panic")
}

type TemplaterTest struct {
	RootPath string
}

func (t *TemplaterTest) Render(name string, view interface{}, response http.ResponseWriter) (interface{}, error) {
	_ = name
	_ = response
	_ = t
	fmt.Println(view)
	return nil, nil
}
