package entity

import (
	"time"

	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
)

// Contact Contact
type Contact struct {
	DateAdd     time.Time `db:"date_add"`
	UUID        string    `db:"uuid"     json:"uuid"`
	Name        string    `db:"name"`
	Message     string    `db:"message"`
	EmailString string
	Email       customtype.NullString `db:"email"`
	ID          int                   `db:"id"    json:"id"`
}
