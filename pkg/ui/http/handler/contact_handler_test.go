package handler_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"net/http/httptest"
	"testing"
)

func TestContactHandler(t *testing.T) {
	t.Run("Unit: test ContactHandler success", func(t *testing.T) {
		contactHandler := handler.ContactHandler{
	     Template: &TemplaterTest{},
			CreateContactCommandHandler: command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          &ContactRepositoryTest{},
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger: &LoggerTest{},
			},
		}

		request := httptest.NewRequest("GET", "/", nil)
		contactHandler.IndexHandle(httptest.NewRecorder(), request)
	})
}

type ContactCreatedEventHandlerTest struct {}

func (h *ContactCreatedEventHandlerTest) Handle(event event.Event) {
	fmt.Println(event)
}

type ContactRepositoryTest struct {}

func (r *ContactRepositoryTest) Save(contact model.IContact) {
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