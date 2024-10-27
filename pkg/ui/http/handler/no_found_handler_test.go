package handler_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"
	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/trace/noop"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
	httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
	httpTracerMock.EXPECT().StartRoot(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	t.Run("Unit: test NotFoundHandler success", func(t *testing.T) {
		_ = t
		notFoundHandler := handler.NotFoundHandler{
			Template: &TemplaterTest{},
			Tracer:   httpTracerMock,
		}
		request := httptest.NewRequest("GET", "/not-found", nil)
		notFoundHandler.Handle(httptest.NewRecorder(), request)
	})
	t.Run("Unit: test NotFoundHandler failed", func(t *testing.T) {
		_ = t
		notFoundHandler := handler.NotFoundHandler{
			Template: &TemplaterTestFailed{},
			Tracer:   httpTracerMock,
		}
		request := httptest.NewRequest("GET", "/not-found", nil)

		notFoundHandler.Handle(httptest.NewRecorder(), request)
	})
}

type TemplaterTestFailed struct {
	RootPath string
}

func (t *TemplaterTestFailed) Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error) {
	_ = name
	_ = view
	_ = response
	_ = status
	return nil, errors.New("panic")
}

type TemplaterTest struct {
	RootPath string
}

func (t *TemplaterTest) Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error) {
	_ = name
	_ = response
	_ = status
	fmt.Println(view)
	return nil, nil
}
