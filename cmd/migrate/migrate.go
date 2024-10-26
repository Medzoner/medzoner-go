package main

import (
	wiring "github.com/Medzoner/medzoner-go/pkg/infra/dependency"
	"log"
)

func main() {
	mg, err := wiring.InitDbMigration()
	if err != nil {
		panic(err)
	}
	mg.DbInstance.CreateDatabase(mg.DbInstance.GetDatabaseName())
	err = mg.MigrateUp()
	if err != nil {
		log.Fatal(err)
	}
}
