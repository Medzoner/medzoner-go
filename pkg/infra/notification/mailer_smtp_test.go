package notification_test

import (
	"context"
	"testing"

	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/infra/notification"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace/noop"
)

func TestSmtp(t *testing.T) {
	t.Run("Unit: test Smtp success", func(t *testing.T) {
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		httpTracerMock.EXPECT().Error(gomock.Any(), gomock.Any()).Return(nil).Times(1)
		mailer := notification.MailerSMTP{RootPath: "./../../..", Tracer: httpTracerMock}
		ctx := context.WithValue(context.Background(), middleware.CorrelationContextKey{}, uuid.New().String())

		_, _ = mailer.Send(ctx, entity.Contact{})
	})
	t.Run("Unit: test Smtp failed with bad RootPath", func(t *testing.T) {
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		mailer := notification.MailerSMTP{RootPath: "", Tracer: httpTracerMock}

		_, _ = mailer.Send(context.Background(), entity.Contact{})
	})
}
