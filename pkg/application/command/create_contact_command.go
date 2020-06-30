package command

import "time"

type CreateContactCommand struct {
	Name    string    `json:"name" validate:"required"`
	Email   string    `json:"email" validate:"required,email"`
	Message string    `json:"message" validate:"required"`
	DateAdd time.Time `json:"dateAdd"`
}

func (c *CreateContactCommandHandler) GetName() string {
	return "CreateContactCommand"
}
