package handler_test

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
	t.Run("Unit: test NotFoundHandler success", func(t *testing.T) {
		notFoundHandler := handler.NotFoundHandler{
			Template: &TemplaterTest{},
			Tracer:   tracer.NewHttpTracer(),
		}
		request := httptest.NewRequest("GET", "/not-found", nil)
		notFoundHandler.Handle(httptest.NewRecorder(), request)
	})
	t.Run("Unit: test NotFoundHandler failed", func(t *testing.T) {
		notFoundHandler := handler.NotFoundHandler{
			Template: &TemplaterTestFailed{},
			Tracer:   tracer.NewHttpTracer(),
		}
		request := httptest.NewRequest("GET", "/not-found", nil)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		notFoundHandler.Handle(httptest.NewRecorder(), request)
	})
}
