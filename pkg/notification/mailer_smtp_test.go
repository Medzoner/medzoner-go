package notification_test

import (
	"context"
	"testing"

	"github.com/Medzoner/gomedz/pkg/logger"
	"github.com/Medzoner/gomedz/pkg/observability"
	"github.com/Medzoner/medzoner-go/internal/entity"
	"github.com/Medzoner/medzoner-go/pkg/notification"
	"github.com/google/uuid"
)

func init() {
	l, err := logger.NewLogger(logger.Config{Level: "debug"})
	if err != nil {
		panic(err)
	}
	_, _ = observability.NewTelemetry(context.Background(), &observability.Config{}, l)
}

func TestSmtp(t *testing.T) {
	t.Run("Unit: test Smtp success", func(t *testing.T) {
		mailer := notification.MailerSMTP{notification.Config{RootPath: "./../../.."}}
		ctx := context.WithValue(context.Background(), notification.CorrelationContextKey{}, uuid.New().String())

		_, _ = mailer.Send(ctx, entity.Contact{})
	})
	t.Run("Unit: test Smtp failed with bad RootPath", func(t *testing.T) {
		mailer := notification.MailerSMTP{notification.Config{RootPath: ""}}

		_, _ = mailer.Send(context.Background(), entity.Contact{})
	})
}
