package mailer

import "context"

// Mailer is an interface that contains method Send that send mail
type Mailer interface {
	Send(ctx context.Context, view interface{}) (bool, error)
}
