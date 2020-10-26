package mailer

//Mailer Mailer
type Mailer interface {
	Send(view interface{}) (bool, error)
}
