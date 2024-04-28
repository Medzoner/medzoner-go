package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/mailersmtp"
	"github.com/Medzoner/medzoner-go/pkg/infra/path"
	"github.com/sarulabs/di"
)

// MailerDefinition MailerDefinition
var MailerDefinition = di.Def{
	Name:  "mailer",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return mailersmtp.NewMailerSMTP(
			ctn.Get("config").(config.IConfig),
			path.RootPath(ctn.Get("config").(config.IConfig).GetRootPath()+"/"),
		), nil
	},
}
