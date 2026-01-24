package mailer

import (
	"context"

	"github.com/Medzoner/medzoner-go/internal/entity"
)

// Mailer is an interface that contains method Send that send mail
//
//go:generate mockgen -destination=../../../../test/mocks/mailer_smtp.go -package=mocks -source=./mailer.go Mailer
type Mailer interface {
	Send(ctx context.Context, view entity.Contact) (bool, error)
}

type MailView struct {
	Subject string
	Body    string
	From    string
	To      []string
}
