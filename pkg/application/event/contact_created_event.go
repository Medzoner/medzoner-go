package event

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
)

type ContactCreatedEvent struct {
	Contact model.IContact
}
