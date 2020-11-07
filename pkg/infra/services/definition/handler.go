package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"github.com/sarulabs/di"
)

//NotFoundHandlerDefinition NotFoundHandlerDefinition
var NotFoundHandlerDefinition = di.Def{
	Name:  "notfound-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return &handler.NotFoundHandler{
			Template: ctn.Get("templater").(templater.Templater),
		}, nil
	},
}

//IndexHandlerDefinition IndexHandlerDefinition
var IndexHandlerDefinition = di.Def{
	Name:  "index-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return &handler.IndexHandler{
			Template: ctn.Get("templater").(templater.Templater),
		}, nil
	},
}

//TechnoHandlerDefinition TechnoHandlerDefinition
var TechnoHandlerDefinition = di.Def{
	Name:  "techno-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return &handler.TechnoHandler{
			Template:               ctn.Get("templater").(templater.Templater),
			ListTechnoQueryHandler: ctn.Get("list-techno-query-handler").(query.ListTechnoQueryHandler),
		}, nil
	},
}

//ContactHandlerDefinition ContactHandlerDefinition
var ContactHandlerDefinition = di.Def{
	Name:  "contact-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return &handler.ContactHandler{
			Template:                    ctn.Get("templater").(templater.Templater),
			CreateContactCommandHandler: ctn.Get("create-contact-command-handler").(command.CreateContactCommandHandler),
			Session:                     ctn.Get("session").(session.Sessioner),
			Validation:                  ctn.Get("validation").(validation.ValidatorAdapter),
		}, nil
	},
}
