package event

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/service/mailer"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"go.opentelemetry.io/otel/attribute"
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

// Handle handles event ContactCreatedEvent and send mail to admin
func (c ContactCreatedEventHandler) Handle(ctx context.Context, event Event) error {
	ctx, iSpan := c.Tracer.Start(ctx, "ContactCreatedEventHandler.Handle")
	defer func() {
		iSpan.End()
	}()
	correlationID := middleware.GetCorrelationID(ctx)
	iSpan.SetAttributes(attribute.String("correlation_id", correlationID))

	switch event.(type) {
	case ContactCreatedEvent:
		_, err := c.Mailer.Send(ctx, event.GetModel())
		if err != nil {
			c.Logger.Error(fmt.Sprintf("Error during send mail: %s", err))
			iSpan.RecordError(err)
			return err
		}
		c.Logger.Log("Mail was send.")
	default:
		c.Logger.Error("Error bad event type.")
	}

	return nil
}
