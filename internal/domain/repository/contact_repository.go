//go:generate mockgen -destination=../../../test/mocks/contact_repository.go -package=domainRepositoryMock -source=./contact_repository.go
package repository

import (
	"context"

	"github.com/Medzoner/medzoner-go/pkg/infra/entity"
)

// ContactRepository ContactRepository
type ContactRepository interface {
	Save(ctx context.Context, contact entity.Contact) error
}
