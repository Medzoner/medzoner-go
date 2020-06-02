package mailer_smtp

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"
)

type MailerSmtp struct {
	User     string
	Password string
	Host     string
	Port     string
}

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (m *MailerSmtp) Send(view interface{}) (bool, error) {
	auth := smtp.PlainAuth("", m.User, m.Password, m.Host)

	r := NewRequest([]string{m.User}, "Hello Junk!", "Hello, World!")

	pp, _ := os.Getwd()
	err := r.ParseTemplate(pp+"/tmpl/Contact/contactEmail.html", view)

	if err != nil {
		return false, err
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := m.Host + ":" + m.Port

	if err2 := smtp.SendMail(addr, auth, m.User, r.to, msg); err2 != nil {
		return false, err2
	}

	return true, nil
}

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
