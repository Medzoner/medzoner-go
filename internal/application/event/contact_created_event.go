package event

import (
	"github.com/Medzoner/medzoner-go/internal/entity"
)

// ContactCreatedEvent is a struct that implements Event interface and contains model Contact
type ContactCreatedEvent struct {
	Contact entity.Contact
}

// GetModel returns model Contact
func (c ContactCreatedEvent) GetModel() any {
	return c.Contact
}
