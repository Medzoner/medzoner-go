package command_test

import (
	"context"
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
		handler := command.NewCreateContactCommandHandler(&ContactRepositoryTest{}, CreateContactEventHandlerTest{}, loggerTest)

		handler.Handle(context.Background(), createContactCommand)
		assert.Equal(t, loggerTest.LogMessages[0], "Contact was created.")
		assert.Equal(t, handler.GetName(), "CreateContactCommand")
		assert.Equal(t, "", contact.GetEmailString())
	})
}

type LoggerTest struct {
	LogMessages []string
}

func (l *LoggerTest) Log(msg string) {
	l.LogMessages = append(l.LogMessages, msg)
	fmt.Println(msg)
}
func (l *LoggerTest) Error(msg string) {
	l.LogMessages = append(l.LogMessages, msg)
	fmt.Println(msg)
}
func (l LoggerTest) New() (logger.ILogger, error) {
	return &LoggerTest{}, nil
}

type ContactRepositoryTest struct{}

func (r ContactRepositoryTest) Save(ctx context.Context, contact model.IContact) {
	fmt.Println(contact)
}

type ContactTest struct {
	entity.Contact
}

func (*ContactTest) New() model.IContact {
	return &ContactTest{}
}

type CreateContactEventHandlerTest struct{}

func (c CreateContactEventHandlerTest) Handle(ctx context.Context, event event.Event) {
	_ = event
}
