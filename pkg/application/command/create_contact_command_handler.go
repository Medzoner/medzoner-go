package command

import (
	"context"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/factory"
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"time"
)

type CreateContactCommandHandler struct {
	ContactFactory             factory.IContactFactory
	ContactRepository          repository.ContactRepository
	ContactCreatedEventHandler event.EventHandler
	Logger                     logger.ILogger
}

func (h *CreateContactCommandHandler) HandlerName() string {
	return "command.CreateContactCommandHandler"
}

func (h *CreateContactCommandHandler) NewCommand() interface{} {
	return &CreateContactCommand{}
}

func (h *CreateContactCommandHandler) Handle(ctx context.Context, c interface{}) error {
	_ = ctx
	cmd := c.(CreateContactCommand)
	h.handle(cmd)
	return nil
}

func (h *CreateContactCommandHandler) handle(command CreateContactCommand) {
	contact := h.ContactFactory.New()
	contact.
		SetName(command.Name).
		SetMessage(command.Description).
		SetEmail(customtype.NullString{String: command.Email, Valid: true}).
		SetDateAdd(time.Now())

	h.ContactRepository.Save(contact)
	h.Logger.Log("Contact was created.")

	contactCreatedEvent := event.ContactCreatedEvent{Contact: contact}
	h.ContactCreatedEventHandler.Handle(contactCreatedEvent)
}
