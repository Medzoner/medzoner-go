package event_test

import (
	"context"
	"fmt"

	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/domain/customtype"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	mocks "github.com/Medzoner/medzoner-go/test"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/telemetry"

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

		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		httpTelemetryMock.EXPECT().Log(gomock.Any(), gomock.Any()).AnyTimes()
		contactCreatedEvent := event.ContactCreatedEvent{
			Contact: contact,
		}

		mailer := &MailerTest{
			isSend: false,
		}
		handler := event.ContactCreatedEventHandler{
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
		handler := event.NewContactCreatedEventHandler(mailer, httpTelemetryMock)

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
		httpTelemetryMock.EXPECT().ErrorSpan(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
		handler := event.NewContactCreatedEventHandler(mailer, httpTelemetryMock)
		contactCreatedEvent := event.ContactCreatedEvent{
			Contact: contact,
		}
		err := handler.Publish(context.Background(), contactCreatedEvent)
		assert.Equal(t, err, nil)
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
