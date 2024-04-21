package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/infra/captcha"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"github.com/sarulabs/di"
)

// NotFoundHandlerDefinition NotFoundHandlerDefinition
var NotFoundHandlerDefinition = di.Def{
	Name:  "notfound-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return handler.NewNotFoundHandler(
			ctn.Get("templater").(templater.Templater),
			ctn.Get("tracer").(tracer.Tracer),
		), nil
	},
}

// IndexHandlerDefinition IndexHandlerDefinition
var IndexHandlerDefinition = di.Def{
	Name:  "index-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return handler.NewIndexHandler(
			ctn.Get("templater").(templater.Templater),
			ctn.Get("list-techno-query-handler").(query.ListTechnoQueryHandler),
			ctn.Get("config").(config.IConfig).GetRecaptchaSiteKey(),
			ctn.Get("create-contact-command-handler").(command.CreateContactCommandHandler),
			ctn.Get("session").(session.Sessioner),
			ctn.Get("validation").(validation.ValidatorAdapter),
			ctn.Get("captcha").(captcha.RecaptchaAdapter),
			ctn.Get("tracer").(tracer.Tracer),
		), nil
	},
}

// TechnoHandlerDefinition TechnoHandlerDefinition
var TechnoHandlerDefinition = di.Def{
	Name:  "techno-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return handler.NewTechnoHandler(
			ctn.Get("templater").(templater.Templater),
			ctn.Get("list-techno-query-handler").(query.ListTechnoQueryHandler),
			ctn.Get("tracer").(tracer.Tracer),
		), nil
	},
}
