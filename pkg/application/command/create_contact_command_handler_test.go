package command_test

import (
	"context"
	"errors"
	"fmt"
	mocks "github.com/Medzoner/medzoner-go/test"
	"testing"
	"time"

	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/trace/noop"
	"gotest.tools/assert"
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

		mocked := mocks.New(t)
		mocked.HttpTracer.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
		mocked.Mailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
		loggerTest := &LoggerTest{}
		handler := command.NewCreateContactCommandHandler(
			&ContactRepositoryTest{}, event.ContactCreatedEventHandler{Mailer: mocked.Mailer, Tracer: mocked.HttpTracer, Logger: loggerTest}, loggerTest, mocked.HttpTracer,
		)

		err := handler.Handle(context.Background(), createContactCommand)

		assert.Equal(t, err, nil)
		assert.Equal(t, loggerTest.LogMessages[0], "Contact was created.")
	})
	t.Run("Unit: test CreateContactCommandHandler error save db", func(t *testing.T) {
		date := time.Time{}
		createContactCommand := command.CreateContactCommand{
			Name:    "a name",
			Email:   "email@example.com",
			Message: "the message",
			DateAdd: date,
		}

		mocked := mocks.New(t)
		mocked.HttpTracer.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
		mocked.HttpTracer.EXPECT().Error(gomock.Any(), gomock.Any()).Return(errors.New("error")).AnyTimes()
		mocked.ContactRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("error")).Times(1)
		loggerTest := &LoggerTest{}

		handler := command.NewCreateContactCommandHandler(
			mocked.ContactRepository, event.ContactCreatedEventHandler{
				Mailer: mocked.Mailer, Tracer: mocked.HttpTracer, Logger: loggerTest,
			},
			loggerTest,
			mocked.HttpTracer,
		)
		err := handler.Handle(context.Background(), createContactCommand)

		assert.Equal(t, err.Error(), "error")
	})
	t.Run("Unit: test CreateContactCommandHandler error send mail", func(t *testing.T) {
		date := time.Time{}
		createContactCommand := command.CreateContactCommand{
			Name:    "a name",
			Email:   "email@example.com",
			Message: "the message",
			DateAdd: date,
		}

		mocked := mocks.New(t)
		mocked.HttpTracer.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
		mocked.HttpTracer.EXPECT().Error(gomock.Any(), gomock.Any()).Return(errors.New("error")).AnyTimes()
		mocked.Mailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(false, errors.New("error")).Times(1)
		loggerTest := &LoggerTest{}

		handler := command.NewCreateContactCommandHandler(
			&ContactRepositoryTest{}, event.ContactCreatedEventHandler{Mailer: mocked.Mailer, Tracer: mocked.HttpTracer, Logger: loggerTest}, loggerTest, mocked.HttpTracer,
		)
		err := handler.Handle(context.Background(), createContactCommand)

		assert.Equal(t, err.Error(), "error")
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
