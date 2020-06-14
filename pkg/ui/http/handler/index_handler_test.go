package handler_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	t.Run("Unit: test ContactHandler success", func(t *testing.T) {
		indexHandler := handler.IndexHandler{
			Template: &TemplaterTest{},
		}
		request := httptest.NewRequest("GET", "/", nil)
		indexHandler.IndexHandle(httptest.NewRecorder(), request)
	})
}

type TemplaterTest struct {
	RootPath string
}

func (t *TemplaterTest) Render(name string, view interface{}, response http.ResponseWriter, status int) {
	fmt.Println(view)
}
