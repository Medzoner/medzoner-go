package handler_test

import (
	"errors"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/infra/repository"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	t.Run("Unit: test IndexHandler success", func(t *testing.T) {
		indexHandler := handler.IndexHandler{
			Template: &TemplaterTest{},
			ListTechnoQueryHandler: query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					Logger:   &LoggerTest{},
					RootPath: "./../../../../",
				},
			},
		}
		request := httptest.NewRequest("GET", "/", nil)
		indexHandler.IndexHandle(httptest.NewRecorder(), request)
	})
	t.Run("Unit: test IndexHandler failed", func(t *testing.T) {
		indexHandler := handler.IndexHandler{
			Template: &TemplaterTestFailed{},
		}
		request := httptest.NewRequest("GET", "/", nil)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		indexHandler.IndexHandle(httptest.NewRecorder(), request)
	})
}

type TemplaterTestFailed struct {
	RootPath string
}

func (t *TemplaterTestFailed) Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error) {
	return nil, errors.New("panic")
}

type TemplaterTest struct {
	RootPath string
}

func (t *TemplaterTest) Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error) {
	fmt.Println(view)
	return nil, nil
}
