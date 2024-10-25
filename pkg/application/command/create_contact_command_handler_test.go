package command_test

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
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

		loggerTest := &LoggerTest{}
		handler := command.NewCreateContactCommandHandler(&ContactRepositoryTest{}, CreateContactEventHandlerTest{}, loggerTest)

		handler.Handle(context.Background(), createContactCommand)
		assert.Equal(t, loggerTest.LogMessages[0], "Contact was created.")
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

type ContactRepositoryTest struct{}

func (r ContactRepositoryTest) Save(ctx context.Context, contact entity.Contact) error {
	_ = ctx
	fmt.Println(contact)
	return nil
}

type CreateContactEventHandlerTest struct{}

func (c CreateContactEventHandlerTest) Handle(ctx context.Context, event event.Event) {
	_ = ctx
	_ = event
}
