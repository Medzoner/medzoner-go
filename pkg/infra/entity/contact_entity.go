package entity

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"time"
)

type Contact struct {
	CommonModel
	name    string                `db:"name"`
	message string                `db:"message"`
	email   customtype.NullString `db:"email"`
	dateAdd time.Time             `db:"date_add"`
}

func (*Contact) New() model.IContact {
	return &Contact{}
}

func (c *Contact) Name() string {
	return c.name
}

func (c *Contact) SetName(name string) model.IContact {
	c.name = name
	return c
}

func (c *Contact) Message() string {
	return c.name
}

func (c *Contact) SetMessage(message string) model.IContact {
	c.message = message
	return c
}

func (c *Contact) Email() customtype.NullString {
	return c.email
}

func (c *Contact) SetEmail(email customtype.NullString) model.IContact {
	c.email = email
	return c
}

func (c *Contact) DateAdd() time.Time {
	return c.dateAdd
}

func (c *Contact) SetDateAdd(dateAdd time.Time) model.IContact {
	c.dateAdd = dateAdd
	return c
}
