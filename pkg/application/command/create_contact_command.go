package command

import "time"

//CreateContactCommand CreateContactCommand
type CreateContactCommand struct {
	Name    string    `json:"name" validate:"required"`
	Email   string    `json:"email" validate:"required,email"`
	Message string    `json:"message" validate:"required"`
	DateAdd time.Time `json:"dateAdd"`
}

//GetName GetName
func (c *CreateContactCommandHandler) GetName() string {
	return "CreateContactCommand"
}
