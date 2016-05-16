package security

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	model.Token
	*jwt.StandardClaims
}
