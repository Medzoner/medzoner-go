package session_test

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/gorilla/sessions"
	"net/http/httptest"
	"testing"
)

func TestSession(t *testing.T) {
	t.Run("Unit: test Session success", func(t *testing.T) {
		sessioner := &session.SessionerAdapter{
			SessionKey: "test-session",
			Store:      sessions.NewCookieStore([]byte("test-session")),
			Values:     map[interface{}]interface{}{},
		}
		request := httptest.NewRequest("GET", "/", nil)
		sessionerInstance := sessioner.Init(request)
		_ = sessionerInstance.Save(request, httptest.NewRecorder())
		_ = sessionerInstance.GetValue("key")
		sessionerInstance.SetValue("key", "true")
		_ = sessionerInstance.GetValue("key")
	})
	t.Run("Unit: test Session failed", func(t *testing.T) {
		sessioner := &session.SessionerAdapter{
			SessionKey: "failed-test-session",
			Store:      nil,
			Values:     map[interface{}]interface{}{},
		}
		request := httptest.NewRequest("GET", "/", nil)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		_ = sessioner.Init(request)
	})
}
