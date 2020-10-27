package database_test

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestDbMigrationSuccess(t *testing.T) {
	t.Run("Unit: test logger log success", func(t *testing.T) {
		fmt.Println("todo")
	})
}

type DbInstanceTest struct{}

func (d *DbInstanceTest) GetDB() interface{} {
	panic("implement me")
}

func (d *DbInstanceTest) GetDriverInstance() interface{} {
	panic("implement me")
}

func (d *DbInstanceTest) GetConnection() interface{} {
	panic("implement me")
}

func (d *DbInstanceTest) GetDriverName() string {
	panic("implement me")
}

func (d *DbInstanceTest) DbConn(dbDriverName string, dsn string, databaseName string) (db *sqlx.DB) {
	panic("implement me")
}

func (d *DbInstanceTest) CreateDatabase() {
	panic("implement me")
}

func (d *DbInstanceTest) DropDatabase() {
	panic("implement me")
}

func (d *DbInstanceTest) GetDatabaseName() string {
	panic("implement me")
}

func (d *DbInstanceTest) New() *DbInstanceTest {
	return d
}
