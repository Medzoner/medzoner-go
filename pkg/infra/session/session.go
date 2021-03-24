package session

import (
	"github.com/gorilla/sessions"
	"net/http"
)

//Sessioner Sessioner
type Sessioner interface {
	GetValue(name string) interface{}
	Save(r *http.Request, w http.ResponseWriter) error
	SetValue(name string, value string)
	Init(request *http.Request) (Sessioner, error)
}

//SessionerAdapter SessionerAdapter
type SessionerAdapter struct {
	SessionKey      string
	sessionInstance *sessions.Session
	Values          map[interface{}]interface{}
	Store           *sessions.CookieStore
}

//Init Init
func (s SessionerAdapter) Init(request *http.Request) (Sessioner, error) {
	newSesion, err := s.Store.Get(request, s.SessionKey)
	s.sessionInstance = newSesion
	return s, err
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
