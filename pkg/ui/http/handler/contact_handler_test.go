package handler_test

import (
	"errors"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"gotest.tools/assert"
	"net/http"
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
				Logger:                     &LoggerTest{},
			},
			Session:    SessionAdapterTest{},
			Validation: validation.ValidatorAdapter{}.New(),
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
				Logger:                     &LoggerTest{},
			},
			Session:    SessionAdapterTest{},
			Validation: validation.ValidatorAdapter{}.New(),
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/contact", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		request.Form = v
		contactHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, repositoryMock.ContactSaved.GetID(), 0)
		assert.Assert(t, len(repositoryMock.ContactSaved.GetUUID()) > 0)
		assert.Assert(t, len(repositoryMock.ContactSaved.GetDateAdd().String()) > 0)
		assert.Equal(t, repositoryMock.ContactSaved.GetName(), "a name")
		assert.Equal(t, repositoryMock.ContactSaved.GetEmail().String, "email@fake.com")
		assert.Equal(t, repositoryMock.ContactSaved.GetMessage(), "a message")
	})
	t.Run("Unit: test ContactHandler with templater render failed", func(t *testing.T) {
		repositoryMock := &ContactRepositoryTest{}

		contactHandler := handler.ContactHandler{
			Template: &TemplaterRanderFailedTest{},
			CreateContactCommandHandler: command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			Session:    SessionAdapterTest{},
			Validation: validation.ValidatorAdapter{}.New(),
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("Get", "/contact", nil)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		contactHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 500)
	})

	t.Run("Unit: test ContactHandler with form submit failed on struct", func(t *testing.T) {
		repositoryMock := &ContactRepositoryTest{}

		contactHandler := handler.ContactHandler{
			Template: &TemplaterTest{},
			CreateContactCommandHandler: command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			Session:    SessionAdapterTest{},
			Validation: ValidatorFailOnStructTest{}.New(),
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/contact", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		request.Form = v
		contactHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 400)
	})
	t.Run("Unit: test ContactHandler with session save failed when submit and valid", func(t *testing.T) {
		repositoryMock := &ContactRepositoryTest{}

		contactHandler := handler.ContactHandler{
			Template: &TemplaterTest{},
			CreateContactCommandHandler: command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			Session:    SessionAdapterFailOnSaveSessionTest{},
			Validation: validation.ValidatorAdapter{}.New(),
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/contact", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		request.Form = v
		contactHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 500)
	})
	t.Run("Unit: test ContactHandler with form submit failed on struct", func(t *testing.T) {
		repositoryMock := &ContactRepositoryTest{}

		contactHandler := handler.ContactHandler{
			Template: &TemplaterTest{},
			CreateContactCommandHandler: command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			Session:    SessionAdapterTest{},
			Validation: ValidatorFailOnStructTest{}.New(),
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/contact", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		request.Form = v
		contactHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 400)
	})
	t.Run("Unit: test ContactHandler with session save failed when not submit", func(t *testing.T) {
		repositoryMock := &ContactRepositoryTest{}

		contactHandler := handler.ContactHandler{
			Template: &TemplaterTest{},
			CreateContactCommandHandler: command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			Session:    SessionAdapterFailOnSaveSessionTest{},
			Validation: validation.ValidatorAdapter{}.New(),
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("Get", "/contact", nil)
		contactHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 500)
	})
	t.Run("Unit: test ContactHandler with session init failed when not submit", func(t *testing.T) {
		repositoryMock := &ContactRepositoryTest{}

		contactHandler := handler.ContactHandler{
			Template: &TemplaterTest{},
			CreateContactCommandHandler: command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			Session:    SessionAdapterFailOnInitSessionTest{},
			Validation: validation.ValidatorAdapter{}.New(),
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("Get", "/contact", nil)
		contactHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 500)
	})
}

type ContactCreatedEventHandlerTest struct{}

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

type SessionAdapterTest struct{}

func (s SessionAdapterTest) New() session.Sessioner {
	return &SessionAdapterTest{}
}
func (s SessionAdapterTest) Init(request *http.Request) (session.Sessioner, error) {
	_ = request
	return s, nil
}
func (s SessionAdapterTest) Save(r *http.Request, w http.ResponseWriter) error {
	_ = r
	_ = w
	return nil
}
func (s SessionAdapterTest) GetValue(name string) interface{} {
	return name
}
func (s SessionAdapterTest) SetValue(name string, value string) {
	_ = name
	_ = value
}

type ValidatorFailOnStructTest struct{}

func (v ValidatorFailOnStructTest) GetErrors() []validation.CustomError {
	var errs []validation.CustomError
	return errs
}
func (v ValidatorFailOnStructTest) Struct(str interface{}) error {
	_ = str
	var err = errors.New("validation failed")
	return err
}
func (v ValidatorFailOnStructTest) New() validation.MzValidator {
	return &ValidatorFailOnStructTest{}
}

type SessionAdapterFailOnSaveSessionTest struct{}

func (s SessionAdapterFailOnSaveSessionTest) GetValue(name string) interface{} {
	return name
}
func (s SessionAdapterFailOnSaveSessionTest) Save(r *http.Request, w http.ResponseWriter) error {
	_ = r
	_ = w
	return errors.New("error save session")
}
func (s SessionAdapterFailOnSaveSessionTest) SetValue(name string, value string) {
	_ = name
	_ = value
}
func (s SessionAdapterFailOnSaveSessionTest) Init(request *http.Request) (session.Sessioner, error) {
	_ = request
	return s, nil
}
func (s SessionAdapterFailOnSaveSessionTest) New() session.Sessioner {
	return &SessionAdapterFailOnSaveSessionTest{}
}

type SessionAdapterFailOnInitSessionTest struct{}

func (s SessionAdapterFailOnInitSessionTest) GetValue(name string) interface{} {
	return name
}
func (s SessionAdapterFailOnInitSessionTest) Save(r *http.Request, w http.ResponseWriter) error {
	_ = r
	_ = w
	return nil
}
func (s SessionAdapterFailOnInitSessionTest) SetValue(name string, value string) {
	_ = name
	_ = value
}
func (s SessionAdapterFailOnInitSessionTest) Init(request *http.Request) (session.Sessioner, error) {
	_ = request
	return nil, errors.New("error init session")
}
func (s SessionAdapterFailOnInitSessionTest) New() session.Sessioner {
	return &SessionAdapterFailOnInitSessionTest{}
}

type TemplaterRanderFailedTest struct {
	TemplaterTest
}

func (t *TemplaterRanderFailedTest) Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error) {
	_ = name
	_ = view
	_ = response
	_ = status
	return nil, errors.New("error render templater")
}
