package database

import (
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/jmoiron/sqlx"
)

//IDbInstance IDbInstance
type IDbInstance interface {
	New(dbDriverName string, dsn string, databaseName string) (db *sqlx.DB)
	GetConnection() (db *sqlx.DB)
	CreateDatabase(close bool)
	DropDatabase(close bool)
	GetDatabaseName() string
	GetDatabaseDriver() database.Driver
}
