//go:generate mockgen -destination=../../../test/mocks/pkg/domain/repository/contact_repository.go -package=contactMock -source=./contact_repository.go
package repository

import (
	"context"

	"github.com/Medzoner/medzoner-go/pkg/domain/model"
)

// ContactRepository ContactRepository
type ContactRepository interface {
	Save(ctx context.Context, contact model.IContact)
}
