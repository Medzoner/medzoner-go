package security

import (
	"encoding/json"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/security/providers"
	"net/http"
	"strings"
)

type IAuth interface {
	Execute(w http.ResponseWriter, r *http.Request) bool
}

type Auth struct {
	Provider *providers.Provider
}

type Exception model.Exception

func (a *Auth) Execute(w http.ResponseWriter, r *http.Request) bool {
	var header = r.Header.Get("Authorization")
	header = strings.TrimSpace(header)
	if header == "" {
		a.forceBasicAuth(w)
		_ = json.NewEncoder(w).Encode(Exception{Message: "Missing Authorization header"})
		return false
	}

	if strings.HasPrefix(header, "Basic ") == false && strings.HasPrefix(header, "Bearer ") == false {
		a.forceBasicAuth(w)
		return false
	}

	if strings.HasPrefix(header, "Basic ") {
		if a.Provider.Authentications[0].Auth(w, r) {
			return true
		}
	}

	if strings.HasPrefix(header, "Bearer ") {
		if a.Provider.Authentications[1].Auth(w, r) {
			return true
		}
	}
	w.WriteHeader(http.StatusForbidden)
	_ = json.NewEncoder(w).Encode(Exception{Message: "Access Denied."})
	return false
}

func (a *Auth) forceBasicAuth(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
	w.WriteHeader(http.StatusUnauthorized)
}
