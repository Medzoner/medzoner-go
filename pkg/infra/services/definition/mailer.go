package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/mailer_smtp"
	"github.com/sarulabs/di"
	"os"
)

var MailerDefinition = di.Def{
	Name:  "mailer",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		rootPath, _ := os.Getwd()
		return &mailer_smtp.MailerSmtp{
			RootPath: rootPath,
			User:     ctn.Get("config").(config.IConfig).GetMailerUser(),
			Password: ctn.Get("config").(config.IConfig).GetMailerPassword(),
			Host:     "smtp.gmail.com",
			Port:     "587",
		}, nil
	},
}
