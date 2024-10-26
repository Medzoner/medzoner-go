package event

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/service/mailer"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
)

// ContactCreatedEventHandler is a struct that implements EventHandler interface and handle ContactCreatedEvent
type ContactCreatedEventHandler struct {
	Mailer mailer.Mailer
	Logger logger.ILogger
	Tracer tracer.Tracer
}

// NewContactCreatedEventHandler is a function that returns a new ContactCreatedEventHandler
func NewContactCreatedEventHandler(mailer mailer.Mailer, logger logger.ILogger, tracer tracer.Tracer) *ContactCreatedEventHandler {
	return &ContactCreatedEventHandler{
		Mailer: mailer,
		Logger: logger,
		Tracer: tracer,
	}
}

// Publish handles event ContactCreatedEvent and send mail to admin
func (c ContactCreatedEventHandler) Publish(ctx context.Context, event Event) error {
	ctx, iSpan := c.Tracer.Start(ctx, "ContactCreatedEventHandler.Publish")
	defer iSpan.End()

	switch event.(type) {
	case ContactCreatedEvent:
		contactCreated := event.GetModel().(entity.Contact)
		_, err := c.Mailer.Send(ctx, contactCreated)
		if err != nil {
			iSpan.RecordError(err)
			return fmt.Errorf("error during send mail: %w", err)
		}
		c.Logger.Log("Mail was send.")
	default:
		c.Logger.Error("Error bad event type.")
	}

	return nil
}
