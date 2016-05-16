package jwt

import (
	"net/http"
	"strings"
)

type AuthJwtFaker struct {
}

func (a *AuthJwtFaker) Auth(w http.ResponseWriter, r *http.Request) bool {
	_ = w
	var header = r.Header.Get("Authorization")
	header = strings.TrimSpace(header)
	if header == "Bearer fakejwttoken|dupond|ORIGAMIX-USER-JWT1-0000-000000000001|ROLE_ORIGAMI_USER" {
		return true
	}
	return false
}
