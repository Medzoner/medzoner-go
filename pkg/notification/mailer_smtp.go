package notification

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"html/template"
	"net/smtp"
	"strconv"

	"github.com/Medzoner/gomedz/pkg/observability"
	"github.com/Medzoner/medzoner-go/internal/entity"
	"go.opentelemetry.io/otel/attribute"
)

type Config struct {
	RootPath string `env:"ROOT_PATH" envDefault:"./"`
	User     string `env:"USER"     envDefault:"medzoner@xxx.fake"`
	Password string `env:"PASSWORD" envDefault:"xxxxxxxxxxxx"`
	Host     string `env:"HOST"     envDefault:"smtp.gmail.com"`
	Port     string `env:"PORT"     envDefault:"587"`
}

// MailerSMTP MailerSMTP
type MailerSMTP struct {
	Config Config
}

// NewMailerSMTP NewMailerSMTP
func NewMailerSMTP(config Config) *MailerSMTP {
	return &MailerSMTP{
		Config: config,
	}
}

// Request Request
type Request struct {
	subject string
	body    string
	to      []string
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
	_, iSpan := observability.StartSpan(ctx, "MailerSMTP.Send")
	defer func() {
		iSpan.End()
	}()
	correlationID := GetCorrelationID(ctx)
	iSpan.SetAttributes(attribute.String("correlation.id", correlationID))

	req := NewRequest([]string{m.Config.User}, "Message [medzoner.com]", "Hello, World!")
	if err := req.parseTemplate(m.Config.RootPath+"/tmpl/contact/contactEmail.html", view); err != nil {
		iSpan.RecordError(err)
		return false, fmt.Errorf("parse template failed: %w", err)
	}

	auth := smtp.PlainAuth(m.Config.User, m.Config.User, m.Config.Password, m.Config.Host)
	if err := smtp.SendMail(fmt.Sprintf("%s:%s", m.Config.Host, m.Config.Port), auth, m.Config.User, req.to, m.message(view)); err != nil {
		//return false, fmt.Errorf("send mail failed: %w", m.Telemetry.ErrorSpan(iSpan, err))
		return false, fmt.Errorf("send mail failed: %w", err)
	}

	return true, nil
}

// message is a function that returns a message
func (m *MailerSMTP) message(view entity.Contact) []byte {
	r, _ := rand.Read(nil)
	messageID := strconv.FormatInt(int64(r), 10) + "@" + m.Config.Host
	return []byte("From: " + m.Config.User + " <" + m.Config.User + ">" + "\r\n" +
		"To: " + m.Config.User + "\r\n" +
		"Subject: " + "Message de [www.medzoner.com]" + "\r\n\r\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n" +
		"Message-ID: <" + messageID + ">\n\n" +
		view.Message + "\r\n")
}

func (r *Request) parseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return fmt.Errorf("error: %w", err)
	}
	r.body = buf.String()

	return nil
}

type CorrelationContextKey struct{}

func GetCorrelationID(ctx context.Context) string {
	if val := ctx.Value(CorrelationContextKey{}); val != nil {
		if correlationID, ok := val.(string); ok {
			return correlationID
		}
	}
	return ""
}
