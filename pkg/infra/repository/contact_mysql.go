package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	otelTrace "go.opentelemetry.io/otel/trace"
)

// MysqlContactRepository MysqlContactRepository
type MysqlContactRepository struct {
	DbInstance database.DbInstantiator
	Logger     logger.ILogger
	Tracer     tracer.Tracer
}

// NewMysqlContactRepository is a function that returns a new MysqlContactRepository
func NewMysqlContactRepository(dbInstance database.DbInstantiator, logger logger.ILogger, tracer tracer.Tracer) *MysqlContactRepository {
	return &MysqlContactRepository{
		DbInstance: dbInstance,
		Logger:     logger,
		Tracer:     tracer,
	}
}

// Save is a function that saves a contact
func (m *MysqlContactRepository) Save(ctx context.Context, contact entity.Contact) error {
	_, iSpan := m.Tracer.Start(ctx, "MysqlContactRepository.Save")
	defer iSpan.End()

	conn, err := m.DbInstance.GetConnection().Begin()
	if err != nil {
		return errorResponse("error during begin transaction", iSpan, err)
	}
	contact.EmailString = contact.Email.String

	stmt, err := conn.Prepare(`INSERT INTO Contact (name, message, email, date_add, uuid) VALUES (?,?,?,?,?)`)
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			m.Logger.Error(fmt.Sprintln(err))
			iSpan.RecordError(err)
		}
	}(stmt)
	if err != nil {
		return errorResponse("error during prepare transaction", iSpan, err)
	}

	_, err = stmt.Exec(contact.Name, contact.Message, contact.EmailString, contact.DateAdd, contact.UUID)
	if err != nil {
		return errorResponse("error during exec statement", iSpan, err)
	}

	if err = conn.Commit(); err != nil {
		return errorResponse("error during commit transaction", iSpan, err)
	}

	return nil
}

func errorResponse(msg string, iSpan otelTrace.Span, err error) error {
	iSpan.RecordError(err)
	return fmt.Errorf("%s: %w", msg, err)
}
