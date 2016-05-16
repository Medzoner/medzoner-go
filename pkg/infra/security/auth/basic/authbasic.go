package basic

import (
	"net/http"
)

type AuthBasic struct {
	User string
	Pass string
}

func (a *AuthBasic) requireAuth(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
	w.WriteHeader(http.StatusForbidden)
}

func (a *AuthBasic) check(r *http.Request) bool {
	username, password, _ := r.BasicAuth()
	return username == a.User && password == a.Pass
}

func (a *AuthBasic) Auth(w http.ResponseWriter, r *http.Request) bool {
	if a.check(r) == false {
		a.requireAuth(w)
		return false
	}
	return true
}
