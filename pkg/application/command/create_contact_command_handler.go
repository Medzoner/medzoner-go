package command

import (
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/factory"
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"time"
)

type CreateContactCommandHandler struct {
	ContactFactory                    factory.IContactFactory
	ContactRepository          repository.ContactRepository
	ContactCreatedEventHandler event.ContactCreatedEventHandler
}

func (c *CreateContactCommandHandler) Handle(command CreateContactCommand) {
	contact := c.ContactFactory.New()
	contact.SetName(command.Name)
	contact.SetMessage(command.Message)
	contact.SetEmail(customtype.NullString{String: command.Email, Valid: true})
	contact.SetDateAdd(time.Now())

	c.ContactRepository.Save(contact)

	contactCreatedEvent := event.ContactCreatedEvent{Contact: contact}
	c.ContactCreatedEventHandler.Handle(contactCreatedEvent)
}
