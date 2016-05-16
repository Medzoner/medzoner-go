package application

import (
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/sarulabs/di"
)

var CreateContactCommandHandlerDefinition = di.Def{
	Name:  "create-contact-command-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return command.CreateContactCommandHandler{
			ContactFactory:             &entity.Contact{},
			ContactRepository:          ctn.Get("contact-repository").(repository.ContactRepository),
			ContactCreatedEventHandler: ctn.Get("contact-event-created-handler").(event.ContactCreatedEventHandler),
		}, nil
	},
}
