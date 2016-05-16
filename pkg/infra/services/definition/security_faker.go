package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/security/auth"
	"github.com/Medzoner/medzoner-go/pkg/infra/security/auth/basic"
	"github.com/Medzoner/medzoner-go/pkg/infra/security/auth/jwt"
	"github.com/Medzoner/medzoner-go/pkg/infra/security/providers"
	"github.com/sarulabs/di"
)

var ProviderFakerDefinition = di.Def{
	Name:  "provider",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		var p = providers.Provider{
			Authentications: []auth.IAuth{
				&basic.AuthBasic{User: "panda", Pass: "test"},
				&jwt.AuthJwtFaker{},
			},
		}
		return &p, nil
	},
}
