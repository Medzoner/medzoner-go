package handler_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/utils/messager"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"gotest.tools/assert"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestContactHandler(t *testing.T) {
	t.Run("Unit: test ContactHandler success", func(t *testing.T) {
		bus := &CommandBusTest{}

		contactHandler := handler.ContactHandler{
			Template: &TemplaterTest{},
			CommandBus: bus,
		}

		request := httptest.NewRequest("GET", "/contact", nil)
		contactHandler.IndexHandle(httptest.NewRecorder(), request)

		assert.Equal(t, len(bus.MessagePublishedList), 0)
	})

	t.Run("Unit: test ContactHandler with form submit success", func(t *testing.T) {
		bus := &CommandBusTest{}
		contactHandler := handler.ContactHandler{
			Template: &TemplaterTest{},
			CommandBus: bus,
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/contact", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		request.Form = v
		contactHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, len(bus.MessagePublishedList), 1)
	})
}

type ContactRepositoryTest struct {
	ContactSaved model.IContact
}

func (r *ContactRepositoryTest) Save(contact model.IContact) {
	r.ContactSaved = contact
	fmt.Println(contact)
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
func (l LoggerTest) New() logger.ILogger {
	return &LoggerTest{}
}

type ContactTest struct {
	entity.Contact
}

func (*ContactTest) New() model.IContact {
	return &ContactTest{}
}
type CommandBusTest struct {
	Bus *cqrs.CommandBus
	Handlers []cqrs.CommandHandler
	MessagePublishedList []messager.Message
}

func (c *CommandBusTest) NewBus() messager.MessageBus {
	return c
}

func (c *CommandBusTest) Handle(message messager.Message)  {
	c.MessagePublishedList = append(c.MessagePublishedList, message)
	fmt.Sprintln(message)
}
