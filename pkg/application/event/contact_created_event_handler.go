package event

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/service/mailer"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
)

// ContactCreatedEventHandler is a struct that implements EventHandler interface and handle ContactCreatedEvent
type ContactCreatedEventHandler struct {
	Mailer mailer.Mailer
	Logger logger.ILogger
}

// NewContactCreatedEventHandler is a function that returns a new ContactCreatedEventHandler
func NewContactCreatedEventHandler(mailer mailer.Mailer, logger logger.ILogger) ContactCreatedEventHandler {
	return ContactCreatedEventHandler{
		Mailer: mailer,
		Logger: logger,
	}
}

// Handle handles event ContactCreatedEvent and send mail to admin
// @param event interface that contains model Contact and event name ContactCreatedEvent (string)
// @return void
func (c ContactCreatedEventHandler) Handle(event Event) {
	switch event.(type) {
	case ContactCreatedEvent:
		_, _ = c.Mailer.Send(event.GetModel())
		err := c.Logger.Log("Mail was send.")
		if err != nil {
			fmt.Println(err)
		}
	default:
		_ = c.Logger.Error("Error during send mail.")
	}
}
