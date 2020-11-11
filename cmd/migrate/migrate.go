package main

import (
	"github.com/Medzoner/medzoner-go/pkg"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sarulabs/di"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	app := pkg.App{
		RootPath: rootPath,
	}
	builder, _ := di.NewBuilder()
	app.LoadContainer(builder)
	app.Container.Get("database").(*database.DbSQLInstance).CreateDatabase(
		app.Container.Get("config").(config.IConfig).GetDatabaseName(),
	)
	app.Container.Get("db-manager").(*database.DbMigration).MigrateUp()
}
