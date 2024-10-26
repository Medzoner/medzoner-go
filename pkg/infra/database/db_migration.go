package database

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	// hack
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	// hack
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

type DbMigration struct {
	DbInstance   IDbInstance
	RootPath     string
	MigrationDir *string
}

// NewDbMigration is a function that returns a new DbMigration
func NewDbMigration(dbInstance IDbInstance, conf config.Config) DbMigration {
	db := DbMigration{
		DbInstance: dbInstance,
		RootPath:   string(conf.RootPath),
	}
	db.MigrationDir = db.getMigrationDir()
	return db

}

// MigrateUp is a function that migrates the database up
func (d *DbMigration) MigrateUp() error {
	db, err := d.getNewWithDatabaseInstance()
	if err != nil {
		return fmt.Errorf("database instantiate failed: %w", err)
	}
	err = db.Up()
	if err != nil {
		return fmt.Errorf("database migration up failed: %w", err)
	}
	err = d.checkMigrateErrors(err)
	if err != nil {
		return fmt.Errorf("database checkMigrateErrors failed: %w", err)
	}
	log.Println("Database migrated ok: up")

	return nil
}

// MigrateDown is a function that migrates down
func (d *DbMigration) MigrateDown() error {
	db, err := d.getNewWithDatabaseInstance()
	if err != nil {
		return fmt.Errorf("database instantiate failed: %w", err)
	}
	err = db.Down()
	if err != nil {
		return fmt.Errorf("database migration down failed: %w", err)
	}
	err = d.checkMigrateErrors(err)
	if err != nil {
		return fmt.Errorf("database checkMigrateErrors failed: %w", err)
	}

	log.Println("Database migrated ok: down")

	return nil
}

func (d *DbMigration) checkMigrateErrors(err error) error {
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("an error occurred while syncing the database.. %w", err)
	}
	return nil
}

func (d *DbMigration) getNewWithDatabaseInstance() (*migrate.Migrate, error) {
	driver, err := d.DbInstance.GetDatabaseDriver()
	if err != nil {
		return nil, fmt.Errorf("database driver failed... %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", *d.MigrationDir), d.DbInstance.GetDatabaseName(), driver)
	if err != nil {
		return nil, fmt.Errorf("migration failed... %w", err)
	}
	return m, nil
}

func (d *DbMigration) getMigrationDir() *string {
	flag.Parse()
	// var migrationDir = flag.String("migration.files", d.RootPath+"/migrations", "Directory where the migration files are located ?")
	migrationDir := d.RootPath + "/migrations"
	return &migrationDir
}
