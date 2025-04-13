package database

import (
	"fmt"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jmoiron/sqlx"
)

type DbSQLInstance struct {
	Connection   *sqlx.DB
	Dsn          string
	DatabaseName string
	DriverName   string
}

// NewDbSQLInstance is a function that returns a new DbSQLInstance
func NewDbSQLInstance(conf config.Config) *DbSQLInstance {
	d := &DbSQLInstance{
		Dsn:          conf.Database.Dsn,
		DatabaseName: conf.Database.Name,
		DriverName:   conf.Database.Driver,
		Connection:   nil,
	}
	d.Connect()
	return d
}

const dsnOptions = "?multiStatements=true&parseTime=true"

// Connect Connect
func (d *DbSQLInstance) Connect() (db *sqlx.DB) {
	d.Connection = d.openDb(d.Dsn + "/" + d.DatabaseName + dsnOptions)
	return d.Connection
}

// GetConnection GetConnection
func (d *DbSQLInstance) GetConnection() (db *sqlx.DB) {
	return d.Connection
}

// CreateDatabase is a function that creates a database
func (d *DbSQLInstance) CreateDatabase(databaseName string) {
	if d.DriverName == "mysql" {
		dbCreate := d.openDb(d.Dsn + "/" + dsnOptions)
		dbCreate.MustExec("CREATE DATABASE IF NOT EXISTS " + databaseName)
	}
}

// DropDatabase DropDatabase
func (d *DbSQLInstance) DropDatabase(databaseName string) {
	if d.DriverName == "mysql" {
		dbDrop := d.openDb(d.Dsn + "/" + dsnOptions)
		dbDrop.MustExec("DROP DATABASE IF EXISTS " + databaseName)
	}
}

// GetDatabaseName GetDatabaseName
func (d *DbSQLInstance) GetDatabaseName() string {
	return d.DatabaseName
}

// GetDatabaseDriver is a function that returns the database driver
func (d *DbSQLInstance) GetDatabaseDriver() (database.Driver, error) {
	db, err := mysql.WithInstance(d.Connection.DB, &mysql.Config{})
	if err != nil {
		return nil, fmt.Errorf("error getting database driver: %w", err)
	}
	return db, nil
}

func (d *DbSQLInstance) openDb(dsn string) *sqlx.DB {
	dbDrop := sqlx.MustOpen(d.DriverName, dsn)
	return dbDrop
}
