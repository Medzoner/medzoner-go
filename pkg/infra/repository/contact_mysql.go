package repository

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
)

// MysqlContactRepository MysqlContactRepository
type MysqlContactRepository struct {
	DbInstance database.DbInstantiator
	Logger     logger.ILogger
	Tracer     tracer.Tracer
}

// NewMysqlContactRepository NewMysqlContactRepository
func NewMysqlContactRepository(dbInstance database.DbInstantiator, logger logger.ILogger, tracer tracer.Tracer) *MysqlContactRepository {
	return &MysqlContactRepository{
		DbInstance: dbInstance,
		Logger:     logger,
		Tracer:     tracer,
	}
}

// Save Save
func (m *MysqlContactRepository) Save(ctx context.Context, contact entity.Contact) error {
	_, iSpan := m.Tracer.Start(ctx, "MysqlContactRepository.Save")
	defer func() {
		iSpan.End()
	}()

	conn := m.DbInstance.GetConnection().MustBegin()
	contact.EmailString = contact.Email.String

	query := `INSERT INTO Contact (name, message, email, date_add, uuid) VALUES (:name, :message, :emailstring, :date_add, :uuid)`
	res, err := conn.NamedExec(query, contact)
	if err != nil {
		m.Logger.Error(fmt.Sprintln(err))
		iSpan.RecordError(err)
		return fmt.Errorf("error during commit transaction: %w", err)
	}
	if res != nil {
		if err = conn.Commit(); err != nil {
			m.Logger.Error(fmt.Sprintln(err))
			iSpan.RecordError(err)
			return fmt.Errorf("error during commit transaction: %w", err)
		}
	}
	return nil
}
