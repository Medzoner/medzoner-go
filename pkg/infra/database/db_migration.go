package database

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	//hack
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	//hack
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

type DbMigration struct {
	DbInstance   IDbInstance
	RootPath     string
	MigrationDir *string
}

// NewDbMigration NewDbMigration
func NewDbMigration(dbInstance IDbInstance, conf config.IConfig) DbMigration {
	db := DbMigration{
		DbInstance: dbInstance,
		RootPath:   string(conf.GetRootPath()),
	}
	db.MigrationDir = db.getMigrationDir()
	return db

}

// MigrateUp MigrateUp
func (d *DbMigration) MigrateUp() {
	err := d.getNewWithDatabaseInstance().Up()
	d.checkMigrateErrors(err)
	log.Println("Database migrated ok: up")
}

// MigrateDown MigrateDown
func (d *DbMigration) MigrateDown() {
	err := d.getNewWithDatabaseInstance().Down()
	d.checkMigrateErrors(err)
	log.Println("Database migrated ok: down")
}

func (d *DbMigration) checkMigrateErrors(err error) {
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}
}

func (d *DbMigration) getNewWithDatabaseInstance() *migrate.Migrate {
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", *d.MigrationDir), d.DbInstance.GetDatabaseName(), d.DbInstance.GetDatabaseDriver())

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}
	return m
}

func (d *DbMigration) getMigrationDir() *string {
	flag.Parse()
	//var migrationDir = flag.String("migration.files", d.RootPath+"/migrations", "Directory where the migration files are located ?")
	migrationDir := d.RootPath + "/migrations"
	return &migrationDir
}
