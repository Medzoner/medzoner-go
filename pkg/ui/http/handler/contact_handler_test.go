package handler_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"gotest.tools/assert"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestContactHandler(t *testing.T) {
	t.Run("Unit: test ContactHandler success", func(t *testing.T) {
		repositoryMock := &ContactRepositoryTest{}
		contactHandler := handler.ContactHandler{
	     Template: &TemplaterTest{},
			CreateContactCommandHandler: command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger: &LoggerTest{},
			},
		}

		request := httptest.NewRequest("GET", "/contact", nil)
		contactHandler.IndexHandle(httptest.NewRecorder(), request)

		assert.Equal(t, repositoryMock.ContactSaved, nil)
	})

	t.Run("Unit: test ContactHandler with form submit success", func(t *testing.T) {
		repositoryMock := &ContactRepositoryTest{}

		contactHandler := handler.ContactHandler{
			Template: &TemplaterTest{},
			CreateContactCommandHandler: command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger: &LoggerTest{},
			},
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/contact", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		request.Form = v
		contactHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, repositoryMock.ContactSaved.Id(), 0)
		assert.Equal(t, repositoryMock.ContactSaved.Name(), "a name")
		assert.Equal(t, repositoryMock.ContactSaved.Email().String, "email@fake.com")
		assert.Equal(t, repositoryMock.ContactSaved.Message(), "a message")
	})
}

type ContactCreatedEventHandlerTest struct {}

func (h *ContactCreatedEventHandlerTest) Handle(event event.Event) {
	fmt.Println(event)
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
