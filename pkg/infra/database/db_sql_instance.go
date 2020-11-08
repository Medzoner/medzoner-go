package database

import (
	"flag"
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
	defer d.closeDb(true)

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
		d.execQuery(dbCreate, "CREATE DATABASE IF NOT EXISTS "+d.DatabaseName)
		d.closeDb(close)
	}
}

//DropDatabase DropDatabase
func (d *DbSQLInstance) DropDatabase(close bool) {
	if d.DriverName == "mysql" {
		dbDrop := d.openDb(d.Dsn + dsnOptions)
		d.execQuery(dbDrop, "DROP DATABASE IF EXISTS "+d.DatabaseName)
		d.closeDb(close)
	}
}

func (d *DbSQLInstance) execQuery(dbCreate *sqlx.DB, query string) {
	_, err := dbCreate.Exec(query)
	if err != nil {
		log.Fatalf("could not create database... %v", err)
	}
}

func (d *DbSQLInstance) closeDb(close bool) {
	if !close {
		err := d.Connection.Close()
		if err != nil {
			log.Fatalf("Could not close database... %v", err)
		}
	}
}

func (d *DbSQLInstance) openDb(dsn string) *sqlx.DB {
	dbDrop, err := sqlx.Open(d.DriverName, dsn)
	if err != nil {
		log.Fatalf("Could not connect to database... %v", err)
	}
	return dbDrop
}

//GetDatabaseName GetDatabaseName
func (d *DbSQLInstance) GetDatabaseName() string {
	return d.DatabaseName
}
