package adapters

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/jmoiron/sqlx"
)

type MysqlDriver struct {
	DB     *sqlx.DB
	Logger logger.ILogger
}

func (d *MysqlDriver) Connect(dsn string, databaseName string) {
	var sqlDSN = flag.String(d.GetName(), dsn+"/"+databaseName+"?multiStatements=true&parseTime=true", d.GetName()+" DSN")
	c, err := sql.Open(d.GetName(), *sqlDSN)
	if err != nil {
		panic(err.Error())
	}
	orm := sqlx.NewDb(c, d.GetName())
	d.DB = orm
}

func (d *MysqlDriver) ExecuteQuery(query string, body interface{}) {
	tx := d.DB.MustBegin()
	res, err := tx.NamedExec(query, body)
	if res != nil {
		_ = tx.Commit()
	}
	if err != nil {
		_ = d.Logger.Error(fmt.Sprintln(err))
	}
}

func (d *MysqlDriver) GetName() string {
	return "mysql"
}
