package command

import (
	"time"
)

type CreateContactCommand struct {
	Name        string    `json:"name" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	Description string    `json:"message" validate:"required"`
	DateAdd     time.Time `json:"dateAdd"`
}

func (c *CreateContactCommand) GetName() string {
	return "CreateContactCommand"
}
