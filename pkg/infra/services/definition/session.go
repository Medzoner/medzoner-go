package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/gorilla/sessions"
	"github.com/sarulabs/di"
)

//SessionDefinition SessionDefinition
var SessionDefinition = di.Def{
	Name:  "session",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return session.SessionerAdapter{
			SessionKey: "medzoner-sessid",
			Store:      sessions.NewCookieStore([]byte("medzoner-sessid")),
		}, nil
	},
}
