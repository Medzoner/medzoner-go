package database

import "github.com/jmoiron/sqlx"

type IDbInstance interface {
	DbConn(dbDriverName string, dsn string, databaseName string) (db *sqlx.DB)
	GetConnection() (db *sqlx.DB)
	CreateDatabase()
	DropDatabase()
	GetDatabaseName() string
}
