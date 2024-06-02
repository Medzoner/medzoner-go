package main

import (
	wiring "github.com/Medzoner/medzoner-go/pkg/infra/dependency"
)

func main() {
	mg, err := wiring.InitDbMigration()
	if err != nil {
		panic(err)
	}
	mg.DbInstance.CreateDatabase(mg.DbInstance.GetDatabaseName())
	mg.MigrateUp()
}
