package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/sarulabs/di"
)

var DatabaseDefinition = di.Def{
	Name:  "database",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		d := database.DbSqlInstance{}
		d.DbConn(
			ctn.Get("config").(config.IConfig).GetDatabaseDriver(),
			ctn.Get("config").(config.IConfig).GetMysqlDsn(),
			ctn.Get("config").(config.IConfig).GetDatabaseName(),
		)
		return &d, nil
	},
}

var DatabaseManagerDefinition = di.Def{
	Name:  "db-manager",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		d := database.DbMigration{
			DbInstance: ctn.Get("database").(database.IDbInstance),
		}
		return &d, nil
	},
}
