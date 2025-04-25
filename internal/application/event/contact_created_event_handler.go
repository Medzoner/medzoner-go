package event

import (
	"context"
	"fmt"

	"github.com/Medzoner/medzoner-go/internal/application/service/mailer"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/telemetry"
)

// ContactCreatedEventHandler is a struct that implements EventHandler interface and handle ContactCreatedEvent
type ContactCreatedEventHandler struct {
	Mailer    mailer.Mailer
	Telemetry telemetry.Telemeter
}

// NewContactCreatedEventHandler is a function that returns a new ContactCreatedEventHandler
func NewContactCreatedEventHandler(mailer mailer.Mailer, tm telemetry.Telemeter) *ContactCreatedEventHandler {
	return &ContactCreatedEventHandler{
		Mailer:    mailer,
		Telemetry: tm,
	}
}

// Publish handles event ContactCreatedEvent and send mail to admin
func (c ContactCreatedEventHandler) Publish(ctx context.Context, event Event) error {
	ctx, iSpan := c.Telemetry.Start(ctx, "ContactCreatedEventHandler.Publish")
	defer iSpan.End()

	switch event.(type) {
	case ContactCreatedEvent:
		if contactCreated, ok := event.GetModel().(entity.Contact); ok {
			if contactCreated.UUID == "" {
				return fmt.Errorf("error during get contact from event: %w", c.Telemetry.ErrorSpan(iSpan, fmt.Errorf("contact UUID is empty")))
			}
			_, err := c.Mailer.Send(ctx, contactCreated)
			if err != nil {
				return fmt.Errorf("error during send mail: %w", c.Telemetry.ErrorSpan(iSpan, err))
			}
			c.Telemetry.Log(ctx, "Mail was send.")
			return nil
		}
		return fmt.Errorf("error during get contact from event: %w", c.Telemetry.ErrorSpan(iSpan, fmt.Errorf("contact is not of type entity.Contact")))
	default:
		c.Telemetry.Error(ctx, "ErrorSpan bad event type.")
	}

	return nil
}
