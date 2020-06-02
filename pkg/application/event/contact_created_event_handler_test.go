package event_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/bmizerany/assert"
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
		SetDateAdd(time.Time{})

	t.Run("Unit: test ContactCreatedEventHandler success", func(t *testing.T) {
		contactCreatedEvent := event.ContactCreatedEvent{
			Contact: contact,
		}

		mailer := &MailerTest{
			isSend: false,
		}
		handler := event.ContactCreatedEventHandler{
			Mailer:  mailer,
			Logger: &LoggerTest{},
		}

		handler.Handle(contactCreatedEvent)
		assert.Equal(t, mailer.isSend, true)
	})
	t.Run("Unit: test ContactCreatedEventHandler failed with bad event", func(t *testing.T) {
		mailer := &MailerTest{
			isSend: false,
		}
		handler := event.ContactCreatedEventHandler{
			Mailer:  mailer,
			Logger: &LoggerTest{},
		}

		handler.Handle(BadEvent{})
		assert.Equal(t, mailer.isSend, false)
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

type BadEvent struct {}
func (b BadEvent) GetName() string {
	return "BadEvent"
}
func (b BadEvent) GetModel() interface{} {
	return BadEvent{}
}
