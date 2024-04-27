package repository

import (
	"context"

	"github.com/Medzoner/medzoner-go/pkg/domain/model"
)

// ContactRepository ContactRepository
type ContactRepository interface {
	Save(ctx context.Context, contact model.IContact)
}
