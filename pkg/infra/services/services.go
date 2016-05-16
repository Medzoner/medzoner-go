package services

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/services/definition"
	"github.com/Medzoner/medzoner-go/pkg/infra/services/definition/application"
	"github.com/sarulabs/di"
)

type Service struct {
	Definitions []di.Def
}

func (s Service) GetDefinitions() []di.Def {
	var services []di.Def

	config := definition.ConfigDependency{}
	config.InitConfig()
	services = append(services, config.GetDefinition())

	services = append(services, definition.LoggerDefinition)
	services = append(services, definition.MailerDefinition)
	services = append(services, definition.DatabaseDefinition)
	services = append(services, definition.DatabaseManagerDefinition)
	services = append(services, definition.WebDefinition)
	services = append(services, definition.RouterDefinition)
	services = append(services, definition.ServerDefinition)
	services = append(services, definition.SecurityDefinition)
	if config.ConfigInstance.GetEnvironment() != "test" && config.ConfigInstance.GetEnvironment() != "test_func" {
		services = append(services, definition.ProviderDefinition)
	}
	services = append(services, definition.ContactRepositoryDefinition)
	services = append(services, definition.TechnoRepositoryDefinition)

	services = append(services, definition.IndexHandlerDefinition)
	services = append(services, definition.TechnoHandlerDefinition)
	services = append(services, definition.ContactHandlerDefinition)

	services = append(services, application.ListTechnoQueryHandlerDefinition)
	services = append(services, application.CreateContactCommandHandlerDefinition)

	services = append(services, application.ContactCreatedEventHandlerDefinition)

	if config.ConfigInstance.GetEnvironment() == "test" || config.ConfigInstance.GetEnvironment() == "test_func" {
		services = append(services, definition.ProviderFakerDefinition)
	}

	return services
}
