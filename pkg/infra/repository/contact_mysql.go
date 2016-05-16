package repository

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/model"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/jmoiron/sqlx"
)

type MysqlContactRepository struct {
	Conn   *sqlx.DB
	Logger logger.ILogger
}

func (m *MysqlContactRepository) Save(contact model.IContact) {
	tx := m.Conn.MustBegin()
	query := `INSERT INTO Contact (name, message, email, date_add) VALUES (:name, :message, :email, :date_add)`
	res, err := tx.NamedExec(query, contact)
	if res != nil {
		_ = tx.Commit()
	}
	if err != nil {
		_ = m.Logger.Error(err)
	}
}
