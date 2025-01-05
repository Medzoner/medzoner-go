package command

import (
	"context"
	"fmt"
	"time"

	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/telemetry"

	"github.com/docker/distribution/uuid"
)

// CreateContactCommandHandler is a struct that implements CommandHandler interface and handle CreateContactCommand
type CreateContactCommandHandler struct {
	ContactRepository          repository.ContactRepository
	ContactCreatedEventHandler event.IEventHandler
	Telemetry                  telemetry.Telemeter
}

// NewCreateContactCommandHandler is a function that returns a new CreateContactCommandHandler
func NewCreateContactCommandHandler(
	contactRepository repository.ContactRepository,
	contactCreatedEventHandler event.IEventHandler,
	tm telemetry.Telemeter,
) CreateContactCommandHandler {
	return CreateContactCommandHandler{
		ContactRepository:          contactRepository,
		ContactCreatedEventHandler: contactCreatedEventHandler,
		Telemetry:                  tm,
	}
}

// Handle handles command CreateContactCommand and create contact in database and send mail to admin with event ContactCreatedEvent
func (c *CreateContactCommandHandler) Handle(ctx context.Context, command CreateContactCommand) error {
	ctx, iSpan := c.Telemetry.Start(ctx, "CreateContactCommandHandler.Publish")
	defer iSpan.End()

	contact := entity.Contact{
		Name:    command.Name,
		Message: command.Message,
		Email:   customtype.NullString{String: command.Email, Valid: true},
		DateAdd: time.Now(),
		UUID:    uuid.UUID{}.String(),
	}
	if err := c.ContactRepository.Save(ctx, contact); err != nil {
		return c.Telemetry.ErrorSpan(iSpan, fmt.Errorf("error during save contact: %w", err))
	}
	c.Telemetry.Log(ctx, "Contact was created.")

	if err := c.ContactCreatedEventHandler.Publish(ctx, event.ContactCreatedEvent{Contact: contact}); err != nil {
		return c.Telemetry.ErrorSpan(iSpan, fmt.Errorf("error during handle event: %w", err))
	}

	return nil
}
