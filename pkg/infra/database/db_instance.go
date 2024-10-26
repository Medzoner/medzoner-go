package database

import (
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/jmoiron/sqlx"
)

type IDbInstance interface {
	GetConnection() (db *sqlx.DB)
	CreateDatabase(databaseName string)
	DropDatabase(databaseName string)
	GetDatabaseName() string
	GetDatabaseDriver() (database.Driver, error)
	Connect() (db *sqlx.DB)
}
