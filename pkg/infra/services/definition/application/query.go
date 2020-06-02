package application

import (
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/sarulabs/di"
)

var ListTechnoQueryHandlerDefinition = di.Def{
	Name:  "list-techno-query-handler",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return query.ListTechnoQueryHandler{
			TechnoRepository: ctn.Get("techno-repository").(repository.TechnoRepository),
		}, nil
	},
}
