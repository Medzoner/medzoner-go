package definition

import (
	command_bus "github.com/Medzoner/medzoner-go/pkg/infra/messagebus"
	"github.com/sarulabs/di"
)

var CommandBusDefinition = di.Def{
	Name:  "message-bus",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return &command_bus.CommandBus{}, nil
	},
}
