package repository

import "github.com/Medzoner/medzoner-go/pkg/domain/model"

//ContactRepository ContactRepository
type ContactRepository interface {
	Save(contact model.IContact)
}
