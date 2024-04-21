package command

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/factory"
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/docker/distribution/uuid"
	"time"
)

// CreateContactCommandHandler is a struct that implements CommandHandler interface and handle CreateContactCommand
type CreateContactCommandHandler struct {
	ContactFactory             factory.IContactFactory
	ContactRepository          repository.ContactRepository
	ContactCreatedEventHandler event.IEventHandler
	Logger                     logger.ILogger
}

// NewCreateContactCommandHandler is a function that returns a new CreateContactCommandHandler
func NewCreateContactCommandHandler(
	contactFactory factory.IContactFactory,
	contactRepository repository.ContactRepository,
	contactCreatedEventHandler event.IEventHandler,
	logger logger.ILogger,
) CreateContactCommandHandler {
	return CreateContactCommandHandler{
		ContactFactory:             contactFactory,
		ContactRepository:          contactRepository,
		ContactCreatedEventHandler: contactCreatedEventHandler,
		Logger:                     logger,
	}
}

// Handle handles command CreateContactCommand and create contact in database and send mail to admin with event ContactCreatedEvent
// @param command CreateContactCommand struct that contains contact data
// @return void
func (c *CreateContactCommandHandler) Handle(command CreateContactCommand) {
	contact := c.ContactFactory.New()
	contact.
		SetName(command.Name).
		SetMessage(command.Message).
		SetEmail(customtype.NullString{String: command.Email, Valid: true}).
		SetDateAdd(time.Now()).
		SetUUID(uuid.UUID{}.String())

	c.ContactRepository.Save(contact)
	err := c.Logger.Log("Contact was created.")
	if err != nil {
		fmt.Println(err)
	}

	contactCreatedEvent := event.ContactCreatedEvent{Contact: contact}
	c.ContactCreatedEventHandler.Handle(contactCreatedEvent)
}
