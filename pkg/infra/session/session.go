package session

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

//Sessioner Sessioner
type Sessioner interface {
	GetValue(name string) interface{}
	New() Sessioner
	Save(r *http.Request, w http.ResponseWriter) error
	SetValue(name string, value string)
	Init(request *http.Request) Sessioner
}

//SessionerAdapter SessionerAdapter
type SessionerAdapter struct {
	SessionKey      string
	sessionInstance *sessions.Session
	Values          map[interface{}]interface{}
	store           *sessions.CookieStore
}

//New New
func (s SessionerAdapter) New() Sessioner {
	return SessionerAdapter{
		SessionKey:      s.SessionKey,
		sessionInstance: s.sessionInstance,
		Values:          s.Values,
		store:           sessions.NewCookieStore([]byte(os.Getenv(s.SessionKey))),
	}
}

//Init Init
func (s SessionerAdapter) Init(request *http.Request) Sessioner {
	newSesion, err := s.store.Get(request, s.SessionKey)
	if err != nil {
		panic(err)
	}
	s.sessionInstance = newSesion
	return s
}

//Save Save
func (s SessionerAdapter) Save(r *http.Request, w http.ResponseWriter) error {
	return s.sessionInstance.Save(r, w)
}

//GetValue GetValue
func (s SessionerAdapter) GetValue(name string) interface{} {
	if len(s.sessionInstance.Values) < 1 {
		var resp interface{}
		return resp
	}
	return s.sessionInstance.Values[name]
}

//SetValue SetValue
func (s SessionerAdapter) SetValue(name string, value string) {
	s.sessionInstance.Values[name] = value
}
