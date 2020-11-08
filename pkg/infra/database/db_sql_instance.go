package database

import (
	"flag"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

//DbSQLInstance DbSQLInstance
type DbSQLInstance struct {
	Connection   *sqlx.DB
	Dsn          string
	DatabaseName string
	DriverName   string
}

const dsnOptions = "?multiStatements=true&parseTime=true"

//DbConn DbConn
func (d *DbSQLInstance) New(dbDriverName string, dsn string, databaseName string) (db *sqlx.DB) {
	d.Dsn = dsn + "/"
	d.DatabaseName = databaseName
	d.DriverName = dbDriverName
	c := d.openDb(*flag.String(dbDriverName, d.Dsn+databaseName+dsnOptions, dbDriverName+" DSN"))
	c.MustExec("USE " + databaseName)

	d.Connection = c

	return c
}

//GetConnection GetConnection
func (d *DbSQLInstance) GetConnection() (db *sqlx.DB) {
	return d.Connection
}

//CreateDatabase CreateDatabase
func (d *DbSQLInstance) CreateDatabase(close bool) {
	if d.DriverName == "mysql" {
		dbCreate := d.openDb(d.Dsn + dsnOptions)
		dbCreate.MustExec("CREATE DATABASE IF NOT EXISTS " + d.DatabaseName)
	}
}

//DropDatabase DropDatabase
func (d *DbSQLInstance) DropDatabase(close bool) {
	if d.DriverName == "mysql" {
		dbDrop := d.openDb(d.Dsn + dsnOptions)
		dbDrop.MustExec("DROP DATABASE IF EXISTS " + d.DatabaseName)
	}
}

func (d *DbSQLInstance) openDb(dsn string) *sqlx.DB {
	dbDrop := sqlx.MustOpen(d.DriverName, dsn)
	return dbDrop
}

//GetDatabaseName GetDatabaseName
func (d *DbSQLInstance) GetDatabaseName() string {
	return d.DatabaseName
}

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
