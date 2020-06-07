package database

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DbSqlInstance struct {
	Connection   *sqlx.DB
	Dsn          string
	DatabaseName string
	DriverName   string
}

func (d *DbSqlInstance) DbConn(dbDriverName string, dsn string, databaseName string) (db *sqlx.DB) {
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

func (d *DbSqlInstance) GetConnection() (db *sqlx.DB) {
	return d.Connection
}

func (d *DbSqlInstance) CreateDatabase() {
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

func (d *DbSqlInstance) DropDatabase() {
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

func (d *DbSqlInstance) GetDatabaseName() string {
	return d.DatabaseName
}
