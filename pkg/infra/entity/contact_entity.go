package entity

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"time"
)

//Contact Contact
type Contact struct {
	ID          int                   `json:"id" db:"id"`
	UUID        string                `json:"uuid" db:"uuid"`
	Name        string                `db:"name"`
	Message     string                `db:"message"`
	Email       customtype.NullString `db:"email"`
	DateAdd     time.Time             `db:"date_add"`
	EmailString string
}

//New New
func (*Contact) New() model.IContact {
	return &Contact{}
}

//GetID GetID
func (c *Contact) GetID() int {
	return c.ID
}

//SetID SetID
func (c *Contact) SetID(id int) model.ICommon {
	c.ID = id
	return c
}

//GetUUID GetUUID
func (c *Contact) GetUUID() string {
	return c.UUID
}

//SetUUID SetUUID
func (c *Contact) SetUUID(uuid string) model.ICommon {
	c.UUID = uuid
	return c
}

//GetName GetName
func (c *Contact) GetName() string {
	return c.Name
}

//SetName SetName
func (c *Contact) SetName(name string) model.IContact {
	c.Name = name
	return c
}

//GetMessage GetMessage
func (c *Contact) GetMessage() string {
	return c.Message
}

//SetMessage SetMessage
func (c *Contact) SetMessage(message string) model.IContact {
	c.Message = message
	return c
}

//GetEmail GetEmail
func (c *Contact) GetEmail() customtype.NullString {
	return c.Email
}

//SetEmail SetEmail
func (c *Contact) SetEmail(email customtype.NullString) model.IContact {
	c.Email = email
	return c
}

//GetEmailString GetEmailString
func (c *Contact) GetEmailString() string {
	return c.Email.String
}

//SetEmailString SetEmailString
func (c *Contact) SetEmailString() model.IContact {
	c.EmailString = c.Email.String
	return c
}

//GetDateAdd GetDateAdd
func (c *Contact) GetDateAdd() time.Time {
	return c.DateAdd
}

//SetDateAdd SetDateAdd
func (c *Contact) SetDateAdd(dateAdd time.Time) model.IContact {
	c.DateAdd = dateAdd
	return c
}
