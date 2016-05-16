package event

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/mailer"
)

type ContactCreatedEventHandler struct {
	Contact model.IContact
	Mailer  *mailer.Mailer
}

func (c *ContactCreatedEventHandler) Handle(event ContactCreatedEvent) {
	_, _ = c.Mailer.Send(event.Contact)
}
