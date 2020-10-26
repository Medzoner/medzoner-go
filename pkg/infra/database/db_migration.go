package database

import (
	"flag"
	"fmt"
	//hack
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	//hack
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

//DbMigration DbMigration
type DbMigration struct {
	DbInstance IDbInstance
	RootPath   string
}

//MigrateUp MigrateUp
func (d *DbMigration) MigrateUp() {
	var migrationDir = flag.String("migration.files", d.RootPath+"/migrations", "Directory where the migration files are located ?")
	var driverInstance database.Driver
	// Run migrations
	if d.DbInstance.GetConnection().DriverName() == "mysql" {
		driver, err := mysql.WithInstance(d.DbInstance.GetConnection().DB, &mysql.Config{})
		if err != nil {
			log.Fatalf("could not start sql migration... %v", err)
		}
		driverInstance = driver
	}
	if d.DbInstance.GetConnection().DriverName() == "sqlite3" {
		driver, err := sqlite3.WithInstance(d.DbInstance.GetConnection().DB, &sqlite3.Config{})
		if err != nil {
			log.Fatalf("could not start sql migration... %v", err)
		}
		driverInstance = driver
	}

	if driverInstance == nil {
		log.Fatalf("driver fail %v", d.DbInstance.GetConnection().DriverName())
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", *migrationDir), d.DbInstance.GetDatabaseName(), driverInstance)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	//_, err = m.Close()
	//if err != nil {
	//	log.Fatalf("could not close database... %v", err)
	//}
	log.Println("Database migrated ok: up")
}

//MigrateDown MigrateDown
func (d *DbMigration) MigrateDown() {
	var migrationDir = flag.String("migration.files", d.RootPath+"/migrations", "Directory where the migration files are located ?")
	// Run migrations
	driverInstance, err := mysql.WithInstance(d.DbInstance.GetConnection().DB, &mysql.Config{})
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", *migrationDir), d.DbInstance.GetDatabaseName(), driverInstance,
	)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	//_, err = m.Close()
	//if err != nil {
	//	log.Fatalf("could not close database... %v", err)
	//}
	log.Println("Database migrated ok: down")
}
