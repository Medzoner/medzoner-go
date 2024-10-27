package event_test

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	mocks "github.com/Medzoner/medzoner-go/test"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"
	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/trace/noop"
	"gotest.tools/assert"
	"reflect"
	"testing"
	"time"
)

func TestContactCreatedEventHandler(t *testing.T) {
	contact := entity.Contact{
		Name:    "a name",
		Email:   customtype.NullString{String: "an email", Valid: true},
		Message: "the message",
		DateAdd: time.Time{},
		ID:      1,
	}

	t.Run("Unit: test ContactCreatedEventHandler success", func(t *testing.T) {

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
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
			Tracer: httpTracerMock,
		}

		err := handler.Publish(context.Background(), contactCreatedEvent)
		assert.Equal(t, err, nil)
		assert.Equal(t, loggerTest.LogMessages[0], "Mail was send.")
		assert.Equal(t, mailer.isSend, true)
	})
	t.Run("Unit: test ContactCreatedEventHandler failed with bad event", func(t *testing.T) {
		mailer := &MailerTest{
			isSend: false,
		}

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		loggerTest := &LoggerTest{}
		handler := event.NewContactCreatedEventHandler(mailer, loggerTest, httpTracerMock)

		err := handler.Publish(context.Background(), BadEvent{})
		assert.Equal(t, err, nil)
		assert.Equal(t, loggerTest.LogMessages[0], "Error bad event type.")
		assert.Equal(t, mailer.isSend, false)
	})
	t.Run("Unit: test ContactCreatedEventHandler failed send mail", func(t *testing.T) {
		mocked := mocks.New(t)
		mailer := mocked.Mailer
		mailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(false, fmt.Errorf("error")).AnyTimes()

		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		httpTracerMock.EXPECT().Error(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
		loggerTest := &LoggerTest{}
		handler := event.NewContactCreatedEventHandler(mailer, loggerTest, httpTracerMock)
		contactCreatedEvent := event.ContactCreatedEvent{
			Contact: contact,
		}
		err := handler.Publish(context.Background(), contactCreatedEvent)
		assert.Equal(t, err, nil)
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

type MailerTest struct {
	User     string
	Password string
	Host     string
	Port     string
	isSend   bool
}

func (m *MailerTest) Send(ctx context.Context, view entity.Contact) (bool, error) {
	_ = ctx
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
