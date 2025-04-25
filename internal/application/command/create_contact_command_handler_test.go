package command_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	command2 "github.com/Medzoner/medzoner-go/internal/application/command"
	"github.com/Medzoner/medzoner-go/internal/application/event"

	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	mocks "github.com/Medzoner/medzoner-go/test"
	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/trace/noop"
	"gotest.tools/assert"
)

func TestCreateContactCommandHandler(t *testing.T) {
	t.Run("Unit: test CreateContactCommandHandler success", func(t *testing.T) {
		date := time.Time{}
		createContactCommand := command2.CreateContactCommand{
			Name:    "a name",
			Email:   "an email",
			Message: "the message",
			DateAdd: date,
		}

		mocked := mocks.New(t)
		mocked.HttpTelemetry.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
		mocked.HttpTelemetry.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
		mocked.HttpTelemetry.EXPECT().Log(gomock.Any(), gomock.Any()).AnyTimes()
		mocked.Mailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
		handler := command2.NewCreateContactCommandHandler(
			&ContactRepositoryTest{}, event.ContactCreatedEventHandler{Mailer: mocked.Mailer, Telemetry: mocked.HttpTelemetry}, mocked.HttpTelemetry,
		)

		err := handler.Handle(context.Background(), createContactCommand)

		assert.Equal(t, err, nil)
	})
	t.Run("Unit: test CreateContactCommandHandler error save db", func(t *testing.T) {
		date := time.Time{}
		createContactCommand := command2.CreateContactCommand{
			Name:    "a name",
			Email:   "email@example.com",
			Message: "the message",
			DateAdd: date,
		}

		mocked := mocks.New(t)
		mocked.HttpTelemetry.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
		mocked.HttpTelemetry.EXPECT().ErrorSpan(gomock.Any(), gomock.Any()).Return(errors.New("error")).AnyTimes()
		mocked.ContactRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("error")).Times(1)

		handler := command2.NewCreateContactCommandHandler(
			mocked.ContactRepository, event.ContactCreatedEventHandler{
				Mailer: mocked.Mailer, Telemetry: mocked.HttpTelemetry,
			},
			mocked.HttpTelemetry,
		)
		err := handler.Handle(context.Background(), createContactCommand)

		assert.Error(t, err, "error during save contact: error")
	})
	t.Run("Unit: test CreateContactCommandHandler error send mail", func(t *testing.T) {
		date := time.Time{}
		createContactCommand := command2.CreateContactCommand{
			Name:    "a name",
			Email:   "email@example.com",
			Message: "the message",
			DateAdd: date,
		}

		mocked := mocks.New(t)
		mocked.HttpTelemetry.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
		mocked.HttpTelemetry.EXPECT().ErrorSpan(gomock.Any(), gomock.Any()).Return(errors.New("error")).AnyTimes()
		mocked.HttpTelemetry.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
		mocked.HttpTelemetry.EXPECT().Log(gomock.Any(), gomock.Any()).AnyTimes()
		mocked.Mailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(false, errors.New("error")).Times(1)

		handler := command2.NewCreateContactCommandHandler(
			&ContactRepositoryTest{}, event.ContactCreatedEventHandler{Mailer: mocked.Mailer, Telemetry: mocked.HttpTelemetry}, mocked.HttpTelemetry,
		)
		err := handler.Handle(context.Background(), createContactCommand)

		assert.Error(t, err, "error during handle event: error")
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
