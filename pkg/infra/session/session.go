package session

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

type Sessioner interface {
	GetValue(name string) interface{}
	New() Sessioner
	Save(r *http.Request, w http.ResponseWriter) error
	SetValue(name string, value string)
	Init(request *http.Request) Sessioner
}

type SessionAdapter struct {
	SessionKey      string
	sessionInstance *sessions.Session
	Values          map[interface{}]interface{}
	store           *sessions.CookieStore
}

func (s SessionAdapter) New() Sessioner {
	return SessionAdapter{
		SessionKey:      s.SessionKey,
		sessionInstance: s.sessionInstance,
		Values:          s.Values,
		store:           sessions.NewCookieStore([]byte(os.Getenv(s.SessionKey))),
	}
}

func (s SessionAdapter) Init(request *http.Request) Sessioner {
	newSesion, err := s.store.Get(request, s.SessionKey)
	if err != nil {
		panic(err)
	}
	s.sessionInstance = newSesion
	return s
}

func (s SessionAdapter) Save(r *http.Request, w http.ResponseWriter) error {
	return s.sessionInstance.Save(r, w)
}

func (s SessionAdapter) GetValue(name string) interface{} {
	if len(s.sessionInstance.Values) < 1 {
		var resp interface{}
		return resp
	}
	return s.sessionInstance.Values[name]
}

func (s SessionAdapter) SetValue(name string, value string) {
	s.sessionInstance.Values[name] = value
}
