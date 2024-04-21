package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/gorilla/sessions"
	"github.com/sarulabs/di"
)

// SessionDefinition SessionDefinition
var SessionDefinition = di.Def{
	Name:  "session",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return session.NewSessionerAdapter(
			"medzoner-sessid",
			sessions.NewCookieStore([]byte("medzoner-sessid")),
		), nil
	},
}
