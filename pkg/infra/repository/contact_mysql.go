package repository

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
)

// MysqlContactRepository MysqlContactRepository
type MysqlContactRepository struct {
	DbInstance database.IDbInstance
	Logger     logger.ILogger
}

// NewMysqlContactRepository NewMysqlContactRepository
func NewMysqlContactRepository(dbInstance database.IDbInstance, logger logger.ILogger) *MysqlContactRepository {
	return &MysqlContactRepository{
		DbInstance: dbInstance,
		Logger:     logger,
	}
}

// Save Save
func (m *MysqlContactRepository) Save(ctx context.Context, contact entity.Contact) error {
	conn := m.DbInstance.GetConnection().MustBegin()
	contact.EmailString = contact.Email.String
	query := `INSERT INTO Contact (name, message, email, date_add, uuid) VALUES (:name, :message, :emailstring, :date_add, :uuid)`
	res, err := m.DbInstance.GetConnection().NamedExec(query, contact)
	if err != nil {
		m.Logger.Error(fmt.Sprintln(err))
		return fmt.Errorf("error during commit transaction: %w", err)
	}
	if res != nil {
		err = conn.Commit()
		if err != nil {
			m.Logger.Error(fmt.Sprintln(err))
			return fmt.Errorf("error during commit transaction: %w", err)
		}
	}
	return nil
}
