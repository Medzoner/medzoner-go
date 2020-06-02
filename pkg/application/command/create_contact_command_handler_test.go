package command_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/bmizerany/assert"
	"testing"
	"time"
)

func TestCreateContactCommandHandler(t *testing.T) {
	t.Run("Unit: test CreateContactCommandHandler success", func(t *testing.T) {
		date := time.Time{}
		createContactCommand := command.CreateContactCommand{
			Name:    "a name",
			Email:   "an email",
			Message: "the message",
			DateAdd: date,
		}

		contact := &ContactTest{}

		handler := command.CreateContactCommandHandler {
			ContactFactory:             contact,
			ContactRepository:          &ContactRepositoryTest{},
			ContactCreatedEventHandler: CreateContactCommandHandlerTest{},
			Logger: &LoggerTest{},
		}

		assert.Equal(t, contact.Name(), "")
		assert.Equal(t, contact.Message(), "")
		assert.Equal(t, contact.Email(), customtype.NullString{String: "", Valid: false})
		assert.Equal(t, contact.DateAdd(), date)
		handler.Handle(createContactCommand)
		//assert.Equal(t, mailer.SendParam, "a name")
	})
}

type LoggerTest struct {
	RootPath string
}
func (l *LoggerTest) Log(msg string) {
	fmt.Println(msg)
}
func (l *LoggerTest) Error(msg string) {
	fmt.Println(msg)
}
func (l LoggerTest) New() logger.ILogger {
	return &LoggerTest{}
}

type ContactRepositoryTest struct {}
func (r ContactRepositoryTest) Save(contact model.IContact) {
	fmt.Println(contact)
}

type ContactTest struct {
	entity.Contact
}
func (*ContactTest) New() model.IContact {
	return &ContactTest{}
}

type CreateContactCommandHandlerTest struct {}

func (c CreateContactCommandHandlerTest) Handle(event event.Event) {
	_ = event
}
