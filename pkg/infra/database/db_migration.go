package database

import (
	"flag"
	"fmt"
	//hack
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
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
	err := d.getNewWithDatabaseInstance().Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("Database migrated ok: up")
}

//MigrateDown MigrateDown
func (d *DbMigration) MigrateDown() {
	err := d.getNewWithDatabaseInstance().Down()
	if  err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("Database migrated ok: down")
}

func (d *DbMigration) getNewWithDatabaseInstance() *migrate.Migrate {
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", *d.getMigrationDir()), d.DbInstance.GetDatabaseName(), d.getMigrateDriver())

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}
	return m
}

func (d *DbMigration) getMigrationDir() *string {
	var migrationDir = flag.String("migration.files", d.RootPath+"/migrations", "Directory where the migration files are located ?")
	return migrationDir
}

func (d *DbMigration) getMigrateDriver() database.Driver {
	driver, err := mysql.WithInstance(d.DbInstance.GetConnection().DB, &mysql.Config{})
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
	}

	if driver == nil {
		log.Fatalf("driver fail %v", d.DbInstance.GetConnection().DriverName())
	}
	return driver
}
