package mailer

// Mailer is an interface that contains method Send that send mail
type Mailer interface {
	Send(view interface{}) (bool, error)
}
