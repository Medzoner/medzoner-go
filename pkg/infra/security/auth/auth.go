package auth

import (
	"net/http"
)

type IAuth interface {
	Auth(w http.ResponseWriter, r *http.Request) bool
}
