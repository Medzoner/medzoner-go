package session_test

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"net/http/httptest"
	"testing"
)

func TestSession(t *testing.T) {
	t.Run("Unit: test Session success", func(t *testing.T) {
		sessioner := &session.SessionAdapter{
			SessionKey: "test-session",
			Values:     map[interface{}]interface{}{},
		}
		request := httptest.NewRequest("GET", "/", nil)
		sessionerInstance := sessioner.New()
		sessionerInstance = sessionerInstance.Init(request)
		_ = sessionerInstance.Save(request, httptest.NewRecorder())
		_ = sessionerInstance.GetValue("key")
		sessionerInstance.SetValue("key", "true")
		_ = sessionerInstance.GetValue("key")
	})
}
