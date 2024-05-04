package main

import (
	wiring "github.com/Medzoner/medzoner-go/pkg/infra/dependency"
)

func main() {
	mg := wiring.InitDbMigration()
	mg.DbInstance.CreateDatabase(mg.DbInstance.GetDatabaseName())
	mg.MigrateUp()
}
