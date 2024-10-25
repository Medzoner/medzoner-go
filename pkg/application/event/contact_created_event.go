package event

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
)

// ContactCreatedEvent is a struct that implements Event interface and contains model Contact
type ContactCreatedEvent struct {
	Contact entity.Contact
}

// GetModel returns model Contact
func (c ContactCreatedEvent) GetModel() interface{} {
	return c.Contact
}
