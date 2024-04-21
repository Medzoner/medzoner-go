package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/sarulabs/di"
)

// DatabaseEntityDefinition DatabaseEntityDefinition
var DatabaseEntityDefinition = di.Def{
	Name:  "database-entity",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		d := database.NewDbSQLInstance(
			nil,
			ctn.Get("config").(config.IConfig).GetMysqlDsn(),
			ctn.Get("config").(config.IConfig).GetDatabaseName(),
			ctn.Get("config").(config.IConfig).GetDatabaseDriver(),
		)
		d.Connect()
		return d, nil
	},
}

// DatabaseDefinition DatabaseDefinition
var DatabaseDefinition = di.Def{
	Name:  "database",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		d := database.NewDbSQLInstance(
			nil,
			ctn.Get("config").(config.IConfig).GetMysqlDsn(),
			ctn.Get("config").(config.IConfig).GetDatabaseName(),
			ctn.Get("config").(config.IConfig).GetDatabaseDriver(),
		)
		d.Connect()
		return d, nil
	},
}

// DatabaseManagerDefinition DatabaseManagerDefinition
var DatabaseManagerDefinition = di.Def{
	Name:  "db-manager",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		d := database.NewDbMigration(
			ctn.Get("database-entity").(database.IDbInstance),
			ctn.Get("config").(config.IConfig).GetRootPath()+"/",
		)
		return d, nil
	},
}
