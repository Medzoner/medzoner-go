package factory

import "github.com/Medzoner/medzoner-go/pkg/domain/model"

type IContactFactory interface {
	New() model.IContact
}
