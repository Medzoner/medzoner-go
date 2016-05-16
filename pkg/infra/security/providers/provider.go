package providers

import "github.com/Medzoner/medzoner-go/pkg/infra/security/auth"

type Provider struct {
	Authentications []auth.IAuth
}
