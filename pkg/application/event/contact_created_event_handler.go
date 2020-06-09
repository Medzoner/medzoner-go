package event

import (
	"github.com/Medzoner/medzoner-go/pkg/application/service/mailer"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
)

type ContactCreatedEventHandler struct {
	Mailer mailer.Mailer
	Logger logger.ILogger
}

func (c ContactCreatedEventHandler) Handle(event Event) {
	switch event.(type) {
	case ContactCreatedEvent:
		_, _ = c.Mailer.Send(event.GetModel())
		c.Logger.Log("Mail was send.")
	default:
		c.Logger.Error("Error during send mail.")
	}
}
