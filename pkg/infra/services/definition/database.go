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
		d := database.DbSQLInstance{
			Connection:   nil,
			Dsn:          ctn.Get("config").(config.IConfig).GetMysqlDsn(),
			DatabaseName: ctn.Get("config").(config.IConfig).GetDatabaseName(),
			DriverName:   ctn.Get("config").(config.IConfig).GetDatabaseDriver(),
		}
		d.Connect()
		return &d, nil
	},
}

// DatabaseDefinition DatabaseDefinition
var DatabaseDefinition = di.Def{
	Name:  "database",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		d := database.DbSQLInstance{
			Connection:   nil,
			Dsn:          ctn.Get("config").(config.IConfig).GetMysqlDsn(),
			DatabaseName: "",
			DriverName:   ctn.Get("config").(config.IConfig).GetDatabaseDriver(),
		}
		d.Connect()
		return &d, nil
	},
}

// DatabaseManagerDefinition DatabaseManagerDefinition
var DatabaseManagerDefinition = di.Def{
	Name:  "db-manager",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		d := database.DbMigration{
			DbInstance: ctn.Get("database-entity").(database.IDbInstance),
			RootPath:   ctn.Get("config").(config.IConfig).GetRootPath() + "/",
		}
		return d.New(), nil
	},
}
