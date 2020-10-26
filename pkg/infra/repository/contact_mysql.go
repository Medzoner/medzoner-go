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
	tx := m.Conn.MustBegin()
	query := `INSERT INTO Contact (name, message, email, date_add, uuid) VALUES (:name, :message, :email, :date_add, :uuid)`
	res, err := tx.NamedExec(query, contact)
	if res != nil {
		_ = tx.Commit()
	}
	if err != nil {
		m.Logger.Error(fmt.Sprintln(err))
	}
}
