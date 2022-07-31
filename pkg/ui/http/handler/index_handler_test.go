package handler_test

import (
	"errors"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	t.Run("Unit: test IndexHandler success", func(t *testing.T) {
		repositoryMock := &ContactRepositoryTest{}
		indexHandler := handler.IndexHandler{
			Template: &TemplaterTest{},
			ListTechnoQueryHandler: query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			CreateContactCommandHandler: command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			Session:    SessionAdapterTest{},
			Validation: validation.ValidatorAdapter{}.New(),
		}
		request := httptest.NewRequest("GET", "/", nil)
		indexHandler.IndexHandle(httptest.NewRecorder(), request)
	})
	t.Run("Unit: test IndexHandler failed", func(t *testing.T) {
		indexHandler := handler.IndexHandler{
			Template: &TemplaterTestFailed{},
		}
		request := httptest.NewRequest("GET", "/", nil)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		indexHandler.IndexHandle(httptest.NewRecorder(), request)
	})
}

type TemplaterTestFailed struct {
	RootPath string
}

func (t *TemplaterTestFailed) Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error) {
	return nil, errors.New("panic")
}

type TemplaterTest struct {
	RootPath string
}

func (t *TemplaterTest) Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error) {
	fmt.Println(view)
	return nil, nil
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
func (l LoggerTest) New() (logger.ILogger, error) {
	return &LoggerTest{}, nil
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
