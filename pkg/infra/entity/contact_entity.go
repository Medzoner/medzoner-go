package entity

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"time"
)

type Contact struct {
	ID      int                   `json:"id" db:"id"`
	UUID    string                `json:"uuid" db:"uuid"`
	Name    string                `db:"name"`
	Message string                `db:"message"`
	Email   customtype.NullString `db:"email"`
	DateAdd time.Time             `db:"date_add"`
}

func (*Contact) New() model.IContact {
	return &Contact{}
}

func (c *Contact) GetID() int {
	return c.ID
}

func (c *Contact) SetID(id int) model.ICommon {
	c.ID = id
	return c
}

func (c *Contact) GetUUID() string {
	return c.UUID
}

func (c *Contact) SetUUID(uuid string) model.ICommon {
	c.UUID = uuid
	return c
}

func (c *Contact) GetName() string {
	return c.Name
}

func (c *Contact) SetName(name string) model.IContact {
	c.Name = name
	return c
}

func (c *Contact) GetMessage() string {
	return c.Message
}

func (c *Contact) SetMessage(message string) model.IContact {
	c.Message = message
	return c
}

func (c *Contact) GetEmail() customtype.NullString {
	return c.Email
}

func (c *Contact) SetEmail(email customtype.NullString) model.IContact {
	c.Email = email
	return c
}

func (c *Contact) GetDateAdd() time.Time {
	return c.DateAdd
}

func (c *Contact) SetDateAdd(dateAdd time.Time) model.IContact {
	c.DateAdd = dateAdd
	return c
}
