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
	ctx, iSpan := c.Tracer.Start(ctx, "CreateContactCommandHandler.Publish")
	defer iSpan.End()

	contact := entity.Contact{
		Name:    command.Name,
		Message: command.Message,
		Email:   customtype.NullString{String: command.Email, Valid: true},
		DateAdd: time.Now(),
		UUID:    uuid.UUID{}.String(),
	}
	if err := c.ContactRepository.Save(ctx, contact); err != nil {
		return c.Tracer.Error(iSpan, fmt.Errorf("error during save contact: %w", err))
	}
	c.Logger.Log("Contact was created.")

	if err := c.ContactCreatedEventHandler.Publish(ctx, event.ContactCreatedEvent{Contact: contact}); err != nil {
		return c.Tracer.Error(iSpan, fmt.Errorf("error during handle event: %w", err))
	}

	return nil
}
