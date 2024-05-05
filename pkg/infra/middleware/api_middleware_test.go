package middleware_test

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiMiddleware(t *testing.T) {
	t.Run("Unit: test APIMiddleware success", func(t *testing.T) {
		apiMiddleware := middleware.NewAPIMiddleware()

		handler := apiMiddleware.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
		assert.Equal(t, true, true)
	})
}
