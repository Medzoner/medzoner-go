package entity

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"time"
)

type Contact struct {
	Id      int                   `json:"id" db:"id"`
	Uuid    string                `json:"uuid" db:"uuid"`
	Name    string                `db:"name"`
	Message string                `db:"message"`
	Email   customtype.NullString `db:"email"`
	DateAdd time.Time             `db:"date_add"`
}

func (*Contact) New() model.IContact {
	return &Contact{}
}

func (c *Contact) GetId() int {
	return c.Id
}

func (c *Contact) SetId(id int) model.ICommon {
	c.Id = id
	return c
}

func (c *Contact) GetUuid() string {
	return c.Uuid
}

func (c *Contact) SetUuid(uuid string) model.ICommon {
	c.Uuid = uuid
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
