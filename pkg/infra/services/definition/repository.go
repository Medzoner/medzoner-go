package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/repository"
	"github.com/sarulabs/di"
)

var ContactRepositoryDefinition = di.Def{
	Name:  "contact-repository",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		r := repository.MysqlContactRepository{
			Logger: ctn.Get("logger").(logger.ILogger),
			Conn:   ctn.Get("database").(*database.DbSqlInstance).Connection,
		}
		return &r, nil
	},
}

var TechnoRepositoryDefinition = di.Def{
	Name:  "techno-repository",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		r := repository.TechnoJsonRepository{
			Logger:   ctn.Get("logger").(logger.ILogger),
			RootPath: ctn.Get("config").(config.IConfig).GetRootPath(),
		}
		return &r, nil
	},
}
