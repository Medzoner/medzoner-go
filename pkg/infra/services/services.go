package services

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/services/definition"
	"github.com/Medzoner/medzoner-go/pkg/infra/services/definition/application"
	"github.com/sarulabs/di"
)

// Service Service
type Service struct {
	Definitions []di.Def
}

var Definitions = []di.Def{
	definition.LoggerDefinition,
	definition.MailerDefinition,
	definition.DatabaseDefinition,
	definition.DatabaseEntityDefinition,
	definition.DatabaseManagerDefinition,
	definition.WebDefinition,
	definition.TemplaterDefinition,
	definition.CaptchaDefinition,
	definition.RouterDefinition,
	definition.SessionDefinition,
	definition.ValidationDefinition,
	definition.ServerDefinition,
	definition.TracerDefinition,
	definition.ContactRepositoryDefinition,
	definition.TechnoRepositoryDefinition,
	definition.NotFoundHandlerDefinition,
	definition.IndexHandlerDefinition,
	definition.TechnoHandlerDefinition,
	application.ListTechnoQueryHandlerDefinition,
	application.CreateContactCommandHandlerDefinition,
	application.ContactCreatedEventHandlerDefinition,
}

// GetDefinitions GetDefinitions
func (s Service) GetDefinitions(rootPath string) []di.Def {
	var services []di.Def

	config := definition.ConfigDependency{}
	config.InitConfig(rootPath)
	services = append(services, config.GetDefinition())

	for _, def := range Definitions {
		services = append(services, def)
	}

	return services
}
