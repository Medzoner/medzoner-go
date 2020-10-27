package repository

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
)

//MysqlContactRepository MysqlContactRepository
type MysqlContactRepository struct {
	DbSQLInstance database.DbSQLInstance
	Logger        logger.ILogger
}

//Save Save
func (m *MysqlContactRepository) Save(contact model.IContact) {
	tx := m.DbSQLInstance.Connection.MustBegin()
	query := `INSERT INTO Contact (name, message, email, date_add, uuid) VALUES (:name, :message, :email, :date_add, :uuid)`
	res, err := tx.NamedExec(query, contact)
	if res != nil {
		_ = tx.Commit()
	}
	if err != nil {
		_ = m.Logger.Error(fmt.Sprintln(err))
	}
}
