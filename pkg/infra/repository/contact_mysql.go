package repository

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
)

//MysqlContactRepository MysqlContactRepository
type MysqlContactRepository struct {
	DbInstance database.IDbInstance
	Logger     logger.ILogger
}

//Save Save
func (m *MysqlContactRepository) Save(contact model.IContact) {
	conn := m.DbInstance.GetConnection().MustBegin()
	contact.SetEmailString()
	query := `INSERT INTO Contact (name, message, email, date_add, uuid) VALUES (:name, :message, :emailstring, :date_add, :uuid)`
	res, err := m.DbInstance.GetConnection().NamedExec(query, contact)
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
