package database

import "github.com/jmoiron/sqlx"

//IDbInstance IDbInstance
type IDbInstance interface {
	DbConn(dbDriverName string, dsn string, databaseName string) (db *sqlx.DB)
	GetConnection() (db *sqlx.DB)
	CreateDatabase(close bool)
	DropDatabase(close bool)
	GetDatabaseName() string
}
