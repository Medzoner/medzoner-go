package command

import "time"

// CreateContactCommand is a struct that contains contact data
type CreateContactCommand struct {
	Name    string    `json:"name" validate:"required"`
	Email   string    `json:"email" validate:"required,email"`
	Message string    `json:"message" validate:"required"`
	DateAdd time.Time `json:"dateAdd"`
}

// GetName returns command name
// @return string
func (c *CreateContactCommandHandler) GetName() string {
	return "CreateContactCommand"
}
