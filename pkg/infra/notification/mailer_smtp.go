package notification

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"html/template"
	"net/smtp"
	"strconv"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"

	"go.opentelemetry.io/otel/attribute"
)

// MailerSMTP MailerSMTP
type MailerSMTP struct {
	RootPath string
	User     string
	Password string
	Host     string
	Port     string
	Tracer   tracer.Tracer
}

// NewMailerSMTP NewMailerSMTP
func NewMailerSMTP(config config.Config, tracer tracer.Tracer) *MailerSMTP {
	return &MailerSMTP{
		RootPath: string(config.RootPath),
		User:     config.Mailer.User,
		Password: config.Mailer.Password,
		Host:     config.Mailer.Host,
		Port:     config.Mailer.Port,
		Tracer:   tracer,
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
func (m *MailerSMTP) Send(ctx context.Context, view entity.Contact) (bool, error) {
	_, iSpan := m.Tracer.Start(ctx, "MailerSMTP.Send")
	defer func() {
		iSpan.End()
	}()
	correlationID := middleware.GetCorrelationID(ctx)
	iSpan.SetAttributes(attribute.String("correlation_id", correlationID))

	req := NewRequest([]string{m.User}, "Message [medzoner.com]", "Hello, World!")
	if err := req.parseTemplate(m.RootPath+"/tmpl/contact/contactEmail.html", view); err != nil {
		iSpan.RecordError(err)
		return false, fmt.Errorf("parse template failed: %w", err)
	}

	auth := smtp.PlainAuth(m.User, m.User, m.Password, m.Host)
	if err := smtp.SendMail(fmt.Sprintf("%s:%s", m.Host, m.Port), auth, m.User, req.to, m.message(view)); err != nil {
		return false, m.Tracer.Error(iSpan, fmt.Errorf("send mail failed: %w", err))
	}

	return true, nil
}

// message is a function that returns a message
func (m *MailerSMTP) message(view entity.Contact) []byte {
	r, _ := rand.Read(nil)
	messageID := strconv.FormatInt(int64(r), 10) + "@" + m.Host
	return []byte("From: " + m.User + " <" + m.User + ">" + "\r\n" +
		"To: " + m.User + "\r\n" +
		"Subject: " + "Message de [www.medzoner.com]" + "\r\n\r\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n" +
		"Message-ID: <" + messageID + ">\n\n" +
		view.Message + "\r\n")
}

func (r *Request) parseTemplate(templateFileName string, data interface{}) error {
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
