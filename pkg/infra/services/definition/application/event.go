package application

import (
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/infra/mailer"
	"github.com/sarulabs/di"
)

var ContactCreatedEventHandlerDefinition = di.Def{
	Name:  "contact-event-created-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return event.ContactCreatedEventHandler{
			Mailer: ctn.Get("mailer").(*mailer.Mailer),
		}, nil
	},
}
