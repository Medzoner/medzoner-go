package database

import (
	"database/sql"
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

//DbConn DbConn
func (d *DbSQLInstance) New(dbDriverName string, dsn string, databaseName string) (db *sqlx.DB) {
	d.DatabaseName = databaseName
	d.DriverName = dbDriverName
	var sqlDSN = flag.String(dbDriverName, dsn+"/"+databaseName+"?multiStatements=true&parseTime=true", d.DriverName+" DSN")
	c := d.openDb(*sqlDSN)
	orm := sqlx.NewDb(c, dbDriverName)
	d.Connection = orm
	d.Dsn = dsn

	return orm
}

//GetConnection GetConnection
func (d *DbSQLInstance) GetConnection() (db *sqlx.DB) {
	return d.Connection
}

//CreateDatabase CreateDatabase
func (d *DbSQLInstance) CreateDatabase(close bool) {
	if d.DriverName == "mysql" {
		dbCreate := d.openDb(d.Dsn + "/?multiStatements=true&parseTime=true")
		d.execQuery(dbCreate, "CREATE DATABASE IF NOT EXISTS "+d.DatabaseName)
		d.closeDb(close)
	}
}

//DropDatabase DropDatabase
func (d *DbSQLInstance) DropDatabase(close bool) {
	if d.DriverName == "mysql" {
		dbDrop := d.openDb(d.Dsn + "/?multiStatements=true&parseTime=true")
		d.execQuery(dbDrop, "DROP DATABASE IF EXISTS "+d.DatabaseName)
		d.closeDb(close)
	}
}

func (d *DbSQLInstance) execQuery(dbCreate *sql.DB, query string) {
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

func (d *DbSQLInstance) openDb(dsn string) *sql.DB {
	dbDrop, err := sql.Open(d.DriverName, dsn)
	if err != nil {
		log.Fatalf("Could not connect to database... %v", err)
	}
	return dbDrop
}

//GetDatabaseName GetDatabaseName
func (d *DbSQLInstance) GetDatabaseName() string {
	return d.DatabaseName
}
