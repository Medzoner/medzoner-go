package session_test

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"net/http/httptest"
	"testing"
)

func TestSession(t *testing.T) {
	t.Run("Unit: test Session success", func(t *testing.T) {
		sess := session.NewSessionerAdapter(session.NewSessionKey())
		request := httptest.NewRequest("GET", "/", nil)
		sessInstance, _ := sess.Init(request)
		_ = sessInstance.Save(request, httptest.NewRecorder())
		_ = sessInstance.GetValue("key")
		sessInstance.SetValue("key", "true")
		_ = sessInstance.GetValue("key")
	})
	t.Run("Unit: test Session failed", func(t *testing.T) {
		sess := &session.SessionerAdapter{
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
		_, _ = sess.Init(request)
	})
}
