package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/sarulabs/di"
)

var SessionDefinition = di.Def{
	Name:  "session",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return session.SessionAdapter{
			SessionKey: "medzoner-sessid",
		}.New(), nil
	},
}
