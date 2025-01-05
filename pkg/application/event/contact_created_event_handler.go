package event

import (
	"context"
	"fmt"

	"github.com/Medzoner/medzoner-go/pkg/application/service/mailer"
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
		contactCreated := event.GetModel().(entity.Contact)
		_, err := c.Mailer.Send(ctx, contactCreated)
		if err != nil {
			return c.Telemetry.ErrorSpan(iSpan, fmt.Errorf("error during send mail: %w", err))
		}
		c.Telemetry.Log(ctx, "Mail was send.")
	default:
		c.Telemetry.Error(ctx, "ErrorSpan bad event type.")
	}

	return nil
}
