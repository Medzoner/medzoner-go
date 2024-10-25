package entity

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"time"
)

// Contact Contact
type Contact struct {
	ID          int                   `json:"id" db:"id"`
	UUID        string                `json:"uuid" db:"uuid"`
	Name        string                `db:"name"`
	Message     string                `db:"message"`
	Email       customtype.NullString `db:"email"`
	DateAdd     time.Time             `db:"date_add"`
	EmailString string
}
