package mailersmtp

import (
	"bytes"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"html/template"
	"net/smtp"
)

// MailerSMTP MailerSMTP
type MailerSMTP struct {
	RootPath string
	User     string
	Password string
	Host     string
	Port     string
}

// NewMailerSMTP NewMailerSMTP
func NewMailerSMTP(config config.Config) *MailerSMTP {
	return &MailerSMTP{
		RootPath: string(config.RootPath),
		User:     config.MailerUser,
		Password: config.MailerPassword,
		Host:     config.MailerHost,
		Port:     config.MailerPort,
	}
}

// Request Request
type Request struct {
	to      []string
	subject string
	body    string
}

// NewRequest NewRequest
func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

// Send is a function that sends an email
func (m *MailerSMTP) Send(view interface{}) (bool, error) {
	auth := smtp.PlainAuth("", m.User, m.Password, m.Host)

	r := NewRequest([]string{m.User}, "Message [medzoner.com]", "Hello, World!")
	if err := r.ParseTemplate(m.RootPath+"/tmpl/contact/contactEmail.html", view); err != nil {
		return false, err
	}

	msg := []byte(fmt.Sprintf(
		"%s%s\n%s",
		fmt.Sprintf("Subject: %s!\n", r.subject),
		"MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n",
		r.body,
	))

	if err := smtp.SendMail(fmt.Sprintf("%s:%s", m.Host, m.Port), auth, m.User, r.to, msg); err != nil {
		return false, err
	}

	return true, nil
}

// ParseTemplate ParseTemplate
func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()

	return nil
}
