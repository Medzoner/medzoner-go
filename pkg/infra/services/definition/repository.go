package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/repository"
	"github.com/sarulabs/di"
)

// ContactRepositoryDefinition ContactRepositoryDefinition
var ContactRepositoryDefinition = di.Def{
	Name:  "contact-repository",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		r := repository.MysqlContactRepository{
			DbInstance: ctn.Get("database-entity").(database.IDbInstance),
			Logger:     ctn.Get("logger").(logger.ILogger),
		}
		return &r, nil
	},
}

// TechnoRepositoryDefinition TechnoRepositoryDefinition
var TechnoRepositoryDefinition = di.Def{
	Name:  "techno-repository",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		r := repository.TechnoJSONRepository{
			Logger:   ctn.Get("logger").(logger.ILogger),
			RootPath: ctn.Get("config").(config.IConfig).GetRootPath(),
		}
		return &r, nil
	},
}
