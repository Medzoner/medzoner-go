//go:generate mockgen -destination=../../../test/mocks/contact_repository.go -package=mocks -source=./contact_repository.go

package repository

import (
	"context"

	"github.com/Medzoner/medzoner-go/internal/entity"
)

type ContactRepository interface {
	Save(ctx context.Context, contact entity.Contact) error
}
