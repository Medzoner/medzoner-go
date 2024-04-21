package handler_test

import (
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/infra/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"gotest.tools/assert"
	"net/http/httptest"
	"testing"
)

func TestTechnoHandler(t *testing.T) {
	t.Run("Unit: test TechnoHandler success", func(t *testing.T) {
		technoHandler := handler.TechnoHandler{
			Template: &TemplaterTest{},
			ListTechnoQueryHandler: query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			Tracer: tracer.NewHttpTracer(),
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/techno", nil)
		technoHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 200)
	})

	t.Run("Unit: test TechnoHandler with templater render failed", func(t *testing.T) {

		technoHandler := handler.TechnoHandler{
			Template: &TemplaterTestFailed{},
			ListTechnoQueryHandler: query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			Tracer: tracer.NewHttpTracer(),
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("Get", "/techno", nil)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		technoHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 500)
	})
}
