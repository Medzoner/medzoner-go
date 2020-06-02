package mailer

type Mailer interface {
	Send(view interface{}) (bool, error)
}
