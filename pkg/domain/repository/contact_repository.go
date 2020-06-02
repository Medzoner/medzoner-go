package repository

import "github.com/Medzoner/medzoner-go/pkg/domain/model"

type ContactRepository interface {
	Save(contact model.IContact)
}
