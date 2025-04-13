package command

import "time"

// CreateContactCommand is a struct that contains contact data
type CreateContactCommand struct {
	DateAdd time.Time `json:"dateAdd"`
	Name    string    `json:"name"    validate:"required"`
	Email   string    `json:"email"   validate:"required,email"`
	Message string    `json:"message" validate:"required"`
}
