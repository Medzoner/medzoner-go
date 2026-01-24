package notification_test

import (
	"context"
	"testing"

	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/infra/notification"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace/noop"
	"github.com/Medzoner/medzoner-go/internal/entity"
)

func TestSmtp(t *testing.T) {
	t.Run("Unit: test Smtp success", func(t *testing.T) {
		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		httpTelemetryMock.EXPECT().ErrorSpan(gomock.Any(), gomock.Any()).Return(nil).Times(1)
		mailer := notification.MailerSMTP{RootPath: "./../../..", Telemetry: httpTelemetryMock}
		ctx := context.WithValue(context.Background(), middleware.CorrelationContextKey{}, uuid.New().String())

		_, _ = mailer.Send(ctx, entity.Contact{})
	})
	t.Run("Unit: test Smtp failed with bad RootPath", func(t *testing.T) {
		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		mailer := notification.MailerSMTP{RootPath: "", Telemetry: httpTelemetryMock}

		_, _ = mailer.Send(context.Background(), entity.Contact{})
	})
}
