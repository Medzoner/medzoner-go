package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/sarulabs/di"
)

//DatabaseDefinition DatabaseDefinition
var DatabaseDefinition = di.Def{
	Name:  "database",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		d := database.DbSQLInstance{}
		d.New(
			ctn.Get("config").(config.IConfig).GetDatabaseDriver(),
			ctn.Get("config").(config.IConfig).GetMysqlDsn(),
			ctn.Get("config").(config.IConfig).GetDatabaseName(),
		)
		return &d, nil
	},
}

//DatabaseManagerDefinition DatabaseManagerDefinition
var DatabaseManagerDefinition = di.Def{
	Name:  "db-manager",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		d := database.DbMigration{
			DbInstance: ctn.Get("database").(database.IDbInstance),
			RootPath:   ctn.Get("config").(config.IConfig).GetRootPath() + "/",
		}
		return d.New(), nil
	},
}
