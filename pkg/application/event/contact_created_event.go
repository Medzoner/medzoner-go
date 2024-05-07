package event

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
)

// ContactCreatedEvent is a struct that implements Event interface and contains model Contact
type ContactCreatedEvent struct {
	Contact model.IContact
}

// GetModel returns model Contact
// @return interface{}
func (c ContactCreatedEvent) GetModel() interface{} {
	return c.Contact
}
