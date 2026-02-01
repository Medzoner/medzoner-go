package mailer

import (
	"context"

	"github.com/Medzoner/gomedz/pkg/notifier"
)

// Mailer is an interface that contains method Send that send mail
//
//go:generate mockgen -destination=../../../../test/mocks/mailer_smtp.go -package=mocks -source=./mailer.go Mailer
type Mailer interface {
	Send(ctx context.Context, view notifier.Message) (bool, error)
}
