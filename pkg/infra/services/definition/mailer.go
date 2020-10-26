package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/mailersmtp"
	"github.com/sarulabs/di"
)

//MailerDefinition MailerDefinition
var MailerDefinition = di.Def{
	Name:  "mailer",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return &mailersmtp.MailerSMTP{
			RootPath: ctn.Get("config").(config.IConfig).GetRootPath() + "/",
			User:     ctn.Get("config").(config.IConfig).GetMailerUser(),
			Password: ctn.Get("config").(config.IConfig).GetMailerPassword(),
			Host:     "smtp.gmail.com",
			Port:     "587",
		}, nil
	},
}
