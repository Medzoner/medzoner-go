package command_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"gotest.tools/assert"
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
		loggerTest := &LoggerTest{}
		handler := command.CreateContactCommandHandler{
			ContactFactory:             contact,
			ContactRepository:          &ContactRepositoryTest{},
			ContactCreatedEventHandler: CreateContactCommandHandlerTest{},
			Logger:                     loggerTest,
		}

		handler.Handle(createContactCommand)
		assert.Equal(t, loggerTest.LogMessages[0], "Contact was created.")
		assert.Equal(t, handler.GetName(), "CreateContactCommand")
	})
}

type LoggerTest struct {
	LogMessages []string
}

func (l *LoggerTest) Log(msg string) error {
	l.LogMessages = append(l.LogMessages, msg)
	fmt.Println(msg)
	return nil
}
func (l *LoggerTest) Error(msg string) error {
	l.LogMessages = append(l.LogMessages, msg)
	fmt.Println(msg)
	return nil
}
func (l LoggerTest) New() logger.ILogger {
	return &LoggerTest{}
}

type ContactRepositoryTest struct{}

func (r ContactRepositoryTest) Save(contact model.IContact) {
	fmt.Println(contact)
}

type ContactTest struct {
	entity.Contact
}

func (*ContactTest) New() model.IContact {
	return &ContactTest{}
}

type CreateContactCommandHandlerTest struct{}

func (c CreateContactCommandHandlerTest) Handle(event event.Event) {
	_ = event
}
