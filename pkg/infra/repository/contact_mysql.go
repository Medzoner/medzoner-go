package repository

import (
	"context"
	"database/sql"
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

	conn, err := m.DbInstance.GetConnection().Begin()
	if err != nil {
		m.Logger.Error(fmt.Sprintln(err))
		iSpan.RecordError(err)
		return fmt.Errorf("error during begin transaction: %w", err)
	}
	contact.EmailString = contact.Email.String

	stmt, err := conn.Prepare(`INSERT INTO Contact (name, message, email, date_add, uuid) VALUES (?,?,?,?,?)`)
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			m.Logger.Error(fmt.Sprintln(err))
			iSpan.RecordError(err)
		}
	}(stmt)
	if err != nil {
		m.Logger.Error(fmt.Sprintln(err))
		iSpan.RecordError(err)
		return fmt.Errorf("error during commit transaction: %w", err)
	}

	_, err = stmt.Exec(contact.Name, contact.Message, contact.EmailString, contact.DateAdd, contact.UUID)
	if err != nil {
		m.Logger.Error(fmt.Sprintln(err))
		iSpan.RecordError(err)
		return fmt.Errorf("error during exec statement: %w", err)
	}

	if err = conn.Commit(); err != nil {
		m.Logger.Error(fmt.Sprintln(err))
		iSpan.RecordError(err)
		return fmt.Errorf("error during commit transaction: %w", err)
	}
	return nil
}
