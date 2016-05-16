package database

import "github.com/jmoiron/sqlx"

type IDbInstance interface {
	DbConn(dbDriver string, dsn string, databaseName string) (db *sqlx.DB)
	GetConnection() (db *sqlx.DB)
	CreateDatabase()
	DropDatabase()
	GetDatabaseName() string
}
