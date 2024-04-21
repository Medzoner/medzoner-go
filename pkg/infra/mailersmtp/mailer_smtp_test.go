package mailersmtp_test

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/mailersmtp"
	"os"
	"testing"
)

func TestSmtp(t *testing.T) {
	t.Run("Unit: test Smtp success", func(t *testing.T) {
		mailer := mailersmtp.MailerSMTP{RootPath: "./../../.."}
		_, _ = mailer.Send(entity.Contact{})
	})
	t.Run("Unit: test Smtp failed with bad RootPath", func(t *testing.T) {
		mailer := mailersmtp.MailerSMTP{RootPath: ""}
		_, _ = mailer.Send(entity.Contact{})
	})
	t.Run("Unit: test Smtp failed with parse error", func(t *testing.T) {
		mailer := mailersmtp.MailerSMTP{RootPath: "./../../../.var/test"}
		_ = os.Chmod(mailer.RootPath+"/tmpl/contact/contactEmail.html", 0000)
		_, _ = mailer.Send(entity.Contact{})
		_ = os.Chmod(mailer.RootPath+"/tmpl/contact/contactEmail.html", 0700)
	})
}
