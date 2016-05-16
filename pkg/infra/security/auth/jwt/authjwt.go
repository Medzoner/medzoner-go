package jwt

import (
	"encoding/json"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/security"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

type AuthJwt struct {
}
type Exception model.Exception

func (a *AuthJwt) Auth(w http.ResponseWriter, r *http.Request) bool {
	var header = r.Header.Get("Authorization")
	header = strings.TrimSpace(header)
	token := &security.Token{}

	_, err := jwt.ParseWithClaims(header, token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(Exception{Message: err.Error()})
		return false
	}
	return true
}
