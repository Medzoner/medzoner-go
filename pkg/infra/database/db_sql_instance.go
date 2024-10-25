package database

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type DbSQLInstance struct {
	Connection   *sqlx.DB
	Dsn          string
	DatabaseName string
	DriverName   string
}

// NewDbSQLInstance NewDbSQLInstance
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

// CreateDatabase CreateDatabase
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

// GetDatabaseDriver GetDatabaseDriver
func (d *DbSQLInstance) GetDatabaseDriver() database.Driver {
	driver, err := mysql.WithInstance(d.Connection.DB, &mysql.Config{})
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
	}

	if driver == nil {
		log.Fatalf("driver fail %v", d.Connection.DriverName())
	}
	return driver
}

func (d *DbSQLInstance) openDb(dsn string) *sqlx.DB {
	dbDrop := sqlx.MustOpen(d.DriverName, dsn)
	return dbDrop
}
