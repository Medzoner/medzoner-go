package event_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	event2 "github.com/Medzoner/medzoner-go/internal/application/event"
	"github.com/Medzoner/medzoner-go/internal/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	mocks "github.com/Medzoner/medzoner-go/test"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks"
	"github.com/golang/mock/gomock"
	
	"go.opentelemetry.io/otel/trace/noop"
	"gotest.tools/assert"
)

func TestContactCreatedEventHandler(t *testing.T) {
	contact := entity.Contact{
		Name:    "a name",
		Email:   customtype.NullString{String: "an email", Valid: true},
		Message: "the message",
		DateAdd: time.Time{},
		ID:      1,
		UUID:    "a uuid",
	}

	t.Run("Unit: test ContactCreatedEventHandler success", func(t *testing.T) {
		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		httpTelemetryMock.EXPECT().Log(gomock.Any(), gomock.Any()).AnyTimes()
		contactCreatedEvent := event2.ContactCreatedEvent{
			Contact: contact,
		}

		mailer := &MailerTest{
			isSend: true,
		}
		handler := event2.ContactCreatedEventHandler{
			Mailer:    mailer,
			Telemetry: httpTelemetryMock,
		}

		err := handler.Publish(context.Background(), contactCreatedEvent)
		assert.Equal(t, err, nil)
		assert.Equal(t, mailer.isSend, true)
	})
	t.Run("Unit: test ContactCreatedEventHandler failed with bad event", func(t *testing.T) {
		mailer := &MailerTest{
			isSend: false,
		}

		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		httpTelemetryMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
		httpTelemetryMock.EXPECT().Log(gomock.Any(), gomock.Any()).AnyTimes()
		handler := event2.NewContactCreatedEventHandler(mailer, httpTelemetryMock)

		err := handler.Publish(context.Background(), BadEvent{})
		assert.Equal(t, err, nil)
		assert.Equal(t, mailer.isSend, false)
	})
	t.Run("Unit: test ContactCreatedEventHandler failed send mail", func(t *testing.T) {
		mocked := mocks.New(t)
		mailer := mocked.Mailer
		mailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(false, fmt.Errorf("error")).AnyTimes()

		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		httpTelemetryMock.EXPECT().ErrorSpan(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error")).AnyTimes()
		handler := event2.NewContactCreatedEventHandler(mailer, httpTelemetryMock)
		contactCreatedEvent := event2.ContactCreatedEvent{
			Contact: contact,
		}
		err := handler.Publish(context.Background(), contactCreatedEvent)
		assert.Error(t, err, "error during send mail: error")
	})
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
	if _, err := fmt.Println(reflect.TypeOf(view)); err != nil {
		return false, fmt.Errorf("error: %w", err)
	}
	return true, nil
}

type BadEvent struct{}

func (b BadEvent) GetName() string {
	return "BadEvent"
}

func (b BadEvent) GetModel() interface{} {
	return BadEvent{}
}
