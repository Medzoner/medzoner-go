package model

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"time"
)

type IContact interface {
	ICommon
	Name() string
	SetName(name string) IContact
	Message() string
	SetMessage(message string) IContact
	Email() customtype.NullString
	SetEmail(email customtype.NullString) IContact
	DateAdd() time.Time
	SetDateAdd(dateAdd time.Time) IContact
}
