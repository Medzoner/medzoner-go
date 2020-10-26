package event_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"gotest.tools/assert"
	"reflect"
	"testing"
	"time"
)

func TestContactCreatedEventHandler(t *testing.T) {
	contact := &ContactTest{}
	contact.
		SetName("a name").
		SetEmail(customtype.NullString{String: "an email", Valid: true}).
		SetMessage("the message").
		SetDateAdd(time.Time{}).
		SetID(1)

	t.Run("Unit: test ContactCreatedEventHandler success", func(t *testing.T) {
		contactCreatedEvent := event.ContactCreatedEvent{
			Contact: contact,
		}

		mailer := &MailerTest{
			isSend: false,
		}
		loggerTest := &LoggerTest{}
		handler := event.ContactCreatedEventHandler{
			Mailer: mailer,
			Logger: loggerTest,
		}

		handler.Handle(contactCreatedEvent)
		assert.Equal(t, loggerTest.LogMessages[0], "Mail was send.")
		assert.Equal(t, mailer.isSend, true)
		assert.Equal(t, contactCreatedEvent.GetName(), "CreateContactCommand")
	})
	t.Run("Unit: test ContactCreatedEventHandler failed with bad event", func(t *testing.T) {
		mailer := &MailerTest{
			isSend: false,
		}
		loggerTest := &LoggerTest{}
		handler := event.ContactCreatedEventHandler{
			Mailer: mailer,
			Logger: loggerTest,
		}

		handler.Handle(BadEvent{})
		assert.Equal(t, loggerTest.LogMessages[0], "Error during send mail.")
		assert.Equal(t, mailer.isSend, false)
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

type ContactTest struct {
	entity.Contact
}

func (*ContactTest) New() model.IContact {
	return &ContactTest{}
}

type MailerTest struct {
	User     string
	Password string
	Host     string
	Port     string
	isSend   bool
}

func (m *MailerTest) Send(view interface{}) (bool, error) {
	m.isSend = true
	_, err := fmt.Println(reflect.TypeOf(view))
	if err != nil {
		m.isSend = false
	}
	return m.isSend, err
}

type BadEvent struct{}

func (b BadEvent) GetName() string {
	return "BadEvent"
}
func (b BadEvent) GetModel() interface{} {
	return BadEvent{}
}
