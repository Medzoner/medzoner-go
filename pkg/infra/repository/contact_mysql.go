package repository

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/jmoiron/sqlx"
)

//MysqlContactRepository MysqlContactRepository
type MysqlContactRepository struct {
	Conn   *sqlx.DB
	Logger logger.ILogger
}

//Save Save
func (m *MysqlContactRepository) Save(contact model.IContact) {
	conn := m.Conn.MustBegin()
	contact.SetEmailString()
	query := `INSERT INTO Contact (name, message, email, date_add, uuid) VALUES (:name, :message, :emailstring, :date_add, :uuid)`
	res, err := conn.NamedExec(query, contact)
	if err != nil {
		_ = m.Logger.Error(fmt.Sprintln(err))
		panic(err)
	}
	if res != nil {
		err = conn.Commit()
		if err != nil {
			_ = m.Logger.Error(fmt.Sprintln(err))
			panic(err)
		}
	}
}
