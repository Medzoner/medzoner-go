package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/messagebus"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sarulabs/di"
)

var CommandBusDefinition = di.Def{
	Name:  "command-bus",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		bus := &messagebus.CommandBus{
			Handlers: []cqrs.CommandHandler{
				ctn.Get("create-contact-command-handler").(cqrs.CommandHandler),
			},
		}
		bus.NewBus()
		return bus, nil
	},
}
