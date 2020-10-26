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
func (d *DbSQLInstance) DbConn(dbDriverName string, dsn string, databaseName string) (db *sqlx.DB) {
	var sqlDSN = flag.String(dbDriverName, dsn+"/"+databaseName+"?multiStatements=true&parseTime=true", d.DriverName+" DSN")
	c, err := sql.Open(dbDriverName, *sqlDSN)
	if err != nil {
		panic(err.Error())
	}
	orm := sqlx.NewDb(c, dbDriverName)
	d.Connection = orm
	d.Dsn = dsn
	d.DatabaseName = databaseName
	d.DriverName = dbDriverName

	return orm
}

//GetConnection GetConnection
func (d *DbSQLInstance) GetConnection() (db *sqlx.DB) {
	return d.Connection
}

//CreateDatabase CreateDatabase
func (d *DbSQLInstance) CreateDatabase() {
	if d.DriverName == "mysql" {
		dbCreate, err := sql.Open(d.DriverName, d.Dsn+"/?multiStatements=true&parseTime=true")
		if err != nil {
			log.Fatalf("could not connect for create database... %v", err)
		}
		_, err = dbCreate.Exec("CREATE DATABASE IF NOT EXISTS " + d.DatabaseName)
		if err != nil {
			log.Fatalf("could not create database... %v", err)
		}
		err = dbCreate.Close()
		if err != nil {
			log.Fatalf("could not close database... %v", err)
		}
	}
}

//DropDatabase DropDatabase
func (d *DbSQLInstance) DropDatabase() {
	if d.DriverName == "mysql" {
		dbDrop, err := sql.Open(d.DriverName, d.Dsn+"/?multiStatements=true&parseTime=true")
		if err != nil {
			log.Fatalf("Could not connect for drop database... %v", err)
		}
		_, err = dbDrop.Exec("DROP DATABASE IF EXISTS " + d.DatabaseName)
		if err != nil {
			log.Fatalf("Could not drop database... %v", err)
		}
		err = dbDrop.Close()
		if err != nil {
			log.Fatalf("Could not close database... %v", err)
		}
	}
}

//GetDatabaseName GetDatabaseName
func (d *DbSQLInstance) GetDatabaseName() string {
	return d.DatabaseName
}
