package main

import (
	"github.com/Medzoner/medzoner-go/pkg"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	app := pkg.App{
		RootPath: rootPath,
	}
	app.LoadContainer()
	app.Container.Get("database").(*database.DbSQLInstance).CreateDatabase(false)
	app.Container.Get("db-manager").(*database.DbMigration).MigrateUp()
}
