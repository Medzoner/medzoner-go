package model

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"time"
)

//IContact IContact
type IContact interface {
	ICommon
	GetName() string
	SetName(name string) IContact
	GetMessage() string
	SetMessage(message string) IContact
	GetEmail() customtype.NullString
	SetEmail(email customtype.NullString) IContact
	GetEmailString() string
	SetEmailString() IContact
	GetDateAdd() time.Time
	SetDateAdd(dateAdd time.Time) IContact
}
