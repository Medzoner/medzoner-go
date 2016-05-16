package repository

import "github.com/Medzoner/medzoner-go/pkg/domain/model"

type ContactRepository interface {
	Save(Contact model.IContact)
}
