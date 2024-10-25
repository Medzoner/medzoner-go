package command

import (
	"context"
	"fmt"
	"time"

	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"

	"github.com/docker/distribution/uuid"
)

// CreateContactCommandHandler is a struct that implements CommandHandler interface and handle CreateContactCommand
type CreateContactCommandHandler struct {
	ContactRepository          repository.ContactRepository
	ContactCreatedEventHandler event.IEventHandler
	Logger                     logger.ILogger
	Tracer                     tracer.Tracer
}

// NewCreateContactCommandHandler is a function that returns a new CreateContactCommandHandler
func NewCreateContactCommandHandler(
	contactRepository repository.ContactRepository,
	contactCreatedEventHandler event.IEventHandler,
	logger logger.ILogger,
	tracer tracer.Tracer,
) CreateContactCommandHandler {
	return CreateContactCommandHandler{
		ContactRepository:          contactRepository,
		ContactCreatedEventHandler: contactCreatedEventHandler,
		Logger:                     logger,
		Tracer:                     tracer,
	}
}

// Handle handles command CreateContactCommand and create contact in database and send mail to admin with event ContactCreatedEvent
func (c *CreateContactCommandHandler) Handle(ctx context.Context, command CreateContactCommand) error {
	_, iSpan := c.Tracer.Start(ctx, fmt.Sprintf("CreateContactCommandHandler.Handle"))
	iSpan.AddEvent("CreateContactCommandHandler.Handle-Event")
	defer func() {
		iSpan.End()
	}()
	contact := entity.Contact{
		Name:    command.Name,
		Message: command.Message,
		Email:   customtype.NullString{String: command.Email, Valid: true},
		DateAdd: time.Now(),
		UUID:    uuid.UUID{}.String(),
	}

	if err := c.ContactRepository.Save(ctx, contact); err != nil {
		c.Logger.Error(fmt.Sprintf("Error during save contact: %s", err))
		return err
	}
	c.Logger.Log("Contact was created.")

	contactCreatedEvent := event.ContactCreatedEvent{Contact: contact}
	if err := c.ContactCreatedEventHandler.Handle(ctx, contactCreatedEvent); err != nil {
		c.Logger.Error(fmt.Sprintf("Error during handle event: %s", err))
		return err
	}

	return nil
}
