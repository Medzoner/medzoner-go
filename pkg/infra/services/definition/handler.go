package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"github.com/sarulabs/di"
)

var IndexHandlerDefinition = di.Def{
	Name:  "index-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return &handler.IndexHandler{
			Template: ctn.Get("templater").(templater.Templater),
		}, nil
	},
}

var TechnoHandlerDefinition = di.Def{
	Name:  "techno-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return &handler.TechnoHandler{
			Template: ctn.Get("templater").(templater.Templater),
			ListTechnoQueryHandler: ctn.Get("list-techno-query-handler").(query.ListTechnoQueryHandler),
		}, nil
	},
}

var ContactHandlerDefinition = di.Def{
	Name:  "contact-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return &handler.ContactHandler{
			Template: ctn.Get("templater").(templater.Templater),
			CreateContactCommandHandler: ctn.Get("create-contact-command-handler").(command.CreateContactCommandHandler),
		}, nil
	},
}
