package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
	"github.com/Medzoner/medzoner-go/pkg/infra/telemetry"
	otelTrace "go.opentelemetry.io/otel/trace"
)

// MysqlContactRepository MysqlContactRepository
type MysqlContactRepository struct {
	DbInstance database.DbInstantiator
	Telemetry  telemetry.Telemeter
}

// NewMysqlContactRepository is a function that returns a new MysqlContactRepository
func NewMysqlContactRepository(dbInstance database.DbInstantiator, tm telemetry.Telemeter) *MysqlContactRepository {
	return &MysqlContactRepository{
		DbInstance: dbInstance,
		Telemetry:  tm,
	}
}

// Save is a function that saves a contact
func (m *MysqlContactRepository) Save(ctx context.Context, contact entity.Contact) error {
	_, iSpan := m.Telemetry.Start(ctx, "MysqlContactRepository.Save")
	defer iSpan.End()

	conn, err := m.DbInstance.GetConnection().Begin()
	if err != nil {
		return errorResponse("error during begin transaction", iSpan, err)
	}
	contact.EmailString = contact.Email.String

	stmt, err := conn.Prepare(`INSERT INTO Contact (name, message, email, date_add, uuid) VALUES (?,?,?,?,?)`)
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			_ = m.Telemetry.ErrorSpan(iSpan, err)
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
