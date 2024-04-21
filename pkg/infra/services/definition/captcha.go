package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/captcha"
	"github.com/sarulabs/di"
)

// CaptchaDefinition CaptchaDefinition
var CaptchaDefinition = di.Def{
	Name:  "captcha",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return captcha.NewRecaptchaAdapter(), nil
	},
}
