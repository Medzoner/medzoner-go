package event

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
)

type ContactCreatedEvent struct {
	Contact model.IContact
}

func (c ContactCreatedEvent) GetModel() interface{} {
	return c.Contact
}

func (c *ContactCreatedEvent) GetName() string {
	return "CreateContactCommand"
}