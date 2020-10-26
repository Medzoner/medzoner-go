package event

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
)

//ContactCreatedEvent ContactCreatedEvent
type ContactCreatedEvent struct {
	Contact model.IContact
}

//GetModel GetModel
func (c ContactCreatedEvent) GetModel() interface{} {
	return c.Contact
}

//GetName GetName
func (c *ContactCreatedEvent) GetName() string {
	return "CreateContactCommand"
}
