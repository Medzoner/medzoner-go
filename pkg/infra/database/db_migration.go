package database

import (
	"flag"
	"fmt"
	//hack
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
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
	d.checkMigrateErrors(err)
	log.Println("Database migrated ok: up")
}

//MigrateDown MigrateDown
func (d *DbMigration) MigrateDown() {
	err := d.getNewWithDatabaseInstance().Down()
	d.checkMigrateErrors(err)
	log.Println("Database migrated ok: down")
}

func (d *DbMigration) checkMigrateErrors(err error) {
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}
}

func (d *DbMigration) getNewWithDatabaseInstance() *migrate.Migrate {
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", *d.getMigrationDir()), d.DbInstance.GetDatabaseName(), d.DbInstance.GetDatabaseDriver())

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}
	return m
}

func (d *DbMigration) getMigrationDir() *string {
	var migrationDir = flag.String("migration.files", d.RootPath+"/migrations", "Directory where the migration files are located ?")
	return migrationDir
}
