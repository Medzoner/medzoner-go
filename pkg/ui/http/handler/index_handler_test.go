package handler_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	mocks "github.com/Medzoner/medzoner-go/test"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"

	metricNoop "go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace/noop"
	"gotest.tools/assert"
)

func TestIndexHandler(t *testing.T) {
	mockedRepository := mocks.New(t)
	t.Run("Unit: test IndexHandler success", func(t *testing.T) {
		repositoryMock := mockedRepository.ContactRepository
		//repositoryMock.EXPECT().Save(gomock.Any(), gomock.Any()).Return()

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(11)
		httpTracerMock.EXPECT().Int64Counter(gomock.Any(), gomock.Any()).Return(metricNoop.Int64Counter{}, nil)
		httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(2)

		indexHandler := handler.NewIndexHandler(
			&TemplaterTest{},
			query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			&config.Config{},
			command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			SessionAdapterTest{},
			validation.ValidatorAdapter{}.New(),
			RecaptchaAdapterTest{},
			httpTracerMock,
			&LoggerTest{},
		)
		request := httptest.NewRequest("GET", "/", nil)
		indexHandler.IndexHandle(httptest.NewRecorder(), request)
	})
	/*	t.Run("Unit: test IndexHandler failed with template error", func(t *testing.T) {
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
	})*/
	t.Run("Unit: test IndexHandler failed with template error on handle", func(t *testing.T) {
		indexHandler := handler.IndexHandler{
			Template: &TemplaterTestFailed{},
			ListTechnoQueryHandler: query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			Session: SessionAdapterTest{},
			Logger:  &LoggerTest{},
		}
		request := httptest.NewRequest("GET", "/", nil)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		indexHandler.IndexHandle(httptest.NewRecorder(), request)
	})
	t.Run("Unit: test IndexHandler failed with session error on Init", func(t *testing.T) {
		indexHandler := handler.IndexHandler{
			Session: &SessionAdapterTestFailed{
				onInit: true,
			},
			Logger: &LoggerTest{},
		}
		request := httptest.NewRequest("GET", "/", nil)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		indexHandler.IndexHandle(httptest.NewRecorder(), request)
	})

	t.Run("Unit: test IndexHandler success", func(t *testing.T) {
		repositoryMock := mockedRepository.ContactRepository

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(11)
		httpTracerMock.EXPECT().Int64Counter(gomock.Any(), gomock.Any()).Return(metricNoop.Int64Counter{}, nil)
		httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(2)

		indexHandler := handler.NewIndexHandler(
			&TemplaterTest{},
			query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			&config.Config{},
			command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			SessionAdapterTest{},
			validation.ValidatorAdapter{}.New(),
			RecaptchaAdapterTest{},
			httpTracerMock,
			&LoggerTest{},
		)

		request := httptest.NewRequest("GET", "/", nil)
		indexHandler.IndexHandle(httptest.NewRecorder(), request)

		// assert.Equal(t, repositoryMock.ContactSaved, nil)
	})

	t.Run("Unit: test IndexHandler with form submit success", func(t *testing.T) {
		repositoryMock := mockedRepository.ContactRepository
		repositoryMock.EXPECT().Save(gomock.Any(), gomock.Any()).Return()

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(11)
		httpTracerMock.EXPECT().Int64Counter(gomock.Any(), gomock.Any()).Return(metricNoop.Int64Counter{}, nil)
		httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(1)

		indexHandler := handler.NewIndexHandler(
			&TemplaterTest{},
			query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			&config.Config{},
			command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			SessionAdapterTest{},
			validation.ValidatorAdapter{}.New(),
			RecaptchaAdapterTest{},
			httpTracerMock,
			&LoggerTest{},
		)

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		v.Set("g-captcha-response", "captcha")
		request.Form = v
		indexHandler.IndexHandle(responseWriter, request)

		// assert.Equal(t, repositoryMock.ContactSaved.GetID(), 0)
		// assert.Assert(t, len(repositoryMock.ContactSaved.GetUUID()) > 0)
		// assert.Assert(t, len(repositoryMock.ContactSaved.GetDateAdd().String()) > 0)
		// assert.Equal(t, repositoryMock.ContactSaved.GetName(), "a name")
		// assert.Equal(t, repositoryMock.ContactSaved.GetEmail().String, "email@fake.com")
		// assert.Equal(t, repositoryMock.ContactSaved.GetMessage(), "a message")
	})
	t.Run("Unit: test IndexHandler with form submit failed on struct", func(t *testing.T) {
		repositoryMock := mockedRepository.ContactRepository

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(11)
		httpTracerMock.EXPECT().Int64Counter(gomock.Any(), gomock.Any()).Return(metricNoop.Int64Counter{}, nil)
		httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(2)

		indexHandler := handler.NewIndexHandler(
			&TemplaterTest{},
			query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			&config.Config{},
			command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			SessionAdapterTest{},
			ValidatorFailOnStructTest{}.New(),
			RecaptchaAdapterTest{},
			httpTracerMock,
			&LoggerTest{},
		)

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		v.Set("g-captcha-response", "captcha")
		request.Form = v
		indexHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 400)
	})
	t.Run("Unit: test IndexHandler with session save failed when submit and valid", func(t *testing.T) {
		repositoryMock := mockedRepository.ContactRepository
		repositoryMock.EXPECT().Save(gomock.Any(), gomock.Any()).Return()

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(11)
		httpTracerMock.EXPECT().Int64Counter(gomock.Any(), gomock.Any()).Return(metricNoop.Int64Counter{}, nil)
		httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(1)

		indexHandler := handler.NewIndexHandler(
			&TemplaterTest{},
			query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			&config.Config{},
			command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			SessionAdapterFailOnSaveSessionTest{},
			validation.ValidatorAdapter{}.New(),
			RecaptchaAdapterTest{},
			httpTracerMock,
			&LoggerTest{},
		)

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		v.Set("g-captcha-response", "captcha")
		request.Form = v
		indexHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 500)
	})
	t.Run("Unit: test IndexHandler with form submit failed on struct", func(t *testing.T) {
		repositoryMock := mockedRepository.ContactRepository

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(11)
		httpTracerMock.EXPECT().Int64Counter(gomock.Any(), gomock.Any()).Return(metricNoop.Int64Counter{}, nil)
		httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(2)

		indexHandler := handler.NewIndexHandler(
			&TemplaterTest{},
			query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			&config.Config{},
			command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			SessionAdapterTest{},
			ValidatorFailOnStructTest{}.New(),
			RecaptchaAdapterTest{},
			httpTracerMock,
			&LoggerTest{},
		)

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		v.Set("g-captcha-response", "captcha")
		request.Form = v
		indexHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 400)
	})
	/*t.Run("Unit: test IndexHandler with session save failed when not submit", func(t *testing.T) {
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
			Session:    SessionAdapterFailOnSaveSessionTest{},
			Validation: validation.ValidatorAdapter{}.New(),
			Recaptcha:  RecaptchaAdapterTest{},
		}

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("Get", "/", nil)
		indexHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 500)
	})*/
	t.Run("Unit: test IndexHandler with session init failed when not submit", func(t *testing.T) {
		repositoryMock := mockedRepository.ContactRepository

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(11)
		httpTracerMock.EXPECT().Int64Counter(gomock.Any(), gomock.Any()).Return(metricNoop.Int64Counter{}, nil)
		httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(1)

		indexHandler := handler.NewIndexHandler(
			&TemplaterTest{},
			query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			&config.Config{},
			command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			SessionAdapterFailOnInitSessionTest{},
			validation.ValidatorAdapter{}.New(),
			RecaptchaAdapterTest{},
			httpTracerMock,
			&LoggerTest{},
		)

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("Get", "/", nil)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		indexHandler.IndexHandle(responseWriter, request)
	})
	t.Run("Unit: test IndexHandler with form submit failed on recaptcha confirm", func(t *testing.T) {
		repositoryMock := mockedRepository.ContactRepository

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(11)
		httpTracerMock.EXPECT().Int64Counter(gomock.Any(), gomock.Any()).Return(metricNoop.Int64Counter{}, nil)
		httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(1)

		indexHandler := handler.NewIndexHandler(
			&TemplaterTest{},
			query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			&config.Config{},
			command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			SessionAdapterTest{},
			validation.ValidatorAdapter{}.New(),
			RecaptchaAdapterTest{
				isFail: true,
			},
			httpTracerMock,
			&LoggerTest{},
		)

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		v.Set("g-captcha-response", "captcha")
		request.Form = v
		indexHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 303)
	})
	t.Run("Unit: test IndexHandler with form submit failed without recaptcha field", func(t *testing.T) {
		repositoryMock := mockedRepository.ContactRepository
		repositoryMock.EXPECT().Save(gomock.Any(), gomock.Any()).Return()

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(11)
		httpTracerMock.EXPECT().Int64Counter(gomock.Any(), gomock.Any()).Return(metricNoop.Int64Counter{}, nil)
		httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(1)

		indexHandler := handler.NewIndexHandler(
			&TemplaterTest{},
			query.ListTechnoQueryHandler{
				TechnoRepository: &repository.TechnoJSONRepository{
					RootPath: "./../../../../",
				},
			},
			&config.Config{},
			command.CreateContactCommandHandler{
				ContactFactory:             &entity.Contact{},
				ContactRepository:          repositoryMock,
				ContactCreatedEventHandler: &ContactCreatedEventHandlerTest{},
				Logger:                     &LoggerTest{},
			},
			SessionAdapterTest{},
			validation.ValidatorAdapter{}.New(),
			RecaptchaAdapterTest{},
			httpTracerMock,
			&LoggerTest{},
		)

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", nil)
		v := url.Values{}
		v.Set("name", "a name")
		v.Set("email", "email@fake.com")
		v.Set("message", "a message")
		request.Form = v
		indexHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 303)
	})
}

type TemplaterTestFailed struct {
	RootPath string
}

func (t *TemplaterTestFailed) Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error) {
	_ = name
	_ = view
	_ = response
	_ = status
	return nil, errors.New("panic")
}

type TemplaterTest struct {
	RootPath string
}

func (t *TemplaterTest) Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error) {
	_ = name
	_ = response
	_ = status
	fmt.Println(view)
	return nil, nil
}

type ContactCreatedEventHandlerTest struct{}

func (h *ContactCreatedEventHandlerTest) Handle(ctx context.Context, event event.Event) {
	fmt.Println(event)
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

type SessionAdapterTestFailed struct {
	SessionAdapterTest
	onInit bool
	onGet  bool
}

func (s SessionAdapterTestFailed) Init(request *http.Request) (session.Sessioner, error) {
	_ = request
	if s.onInit {
		return nil, errors.New("SessionAdapterTestFailed - Init")
	}
	return s, nil
}

func (s SessionAdapterTestFailed) GetValue(name string) interface{} {
	_ = name
	if s.onGet {
		return errors.New("SessionAdapterTestFailed - Init")
	}
	return name
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

type RecaptchaAdapterTest struct {
	isFail bool
}

func (s RecaptchaAdapterTest) Confirm(remoteip, response string) (result bool, err error) {
	_ = remoteip
	_ = response
	if !s.isFail {
		return true, nil
	}
	return false, errors.New("error Confirm Recaptcha")
}
