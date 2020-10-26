package command

import (
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/factory"
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/docker/distribution/uuid"
	"time"
)

type CreateContactCommandHandler struct {
	ContactFactory             factory.IContactFactory
	ContactRepository          repository.ContactRepository
	ContactCreatedEventHandler event.EventHandler
	Logger                     logger.ILogger
}

func (c *CreateContactCommandHandler) Handle(command CreateContactCommand) {
	contact := c.ContactFactory.New()
	contact.
		SetName(command.Name).
		SetMessage(command.Message).
		SetEmail(customtype.NullString{String: command.Email, Valid: true}).
		SetDateAdd(time.Now()).
		SetUUID(uuid.UUID{}.String())

	c.ContactRepository.Save(contact)
	c.Logger.Log("Contact was created.")

	contactCreatedEvent := event.ContactCreatedEvent{Contact: contact}
	c.ContactCreatedEventHandler.Handle(contactCreatedEvent)
}
