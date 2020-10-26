package factory

import "github.com/Medzoner/medzoner-go/pkg/domain/model"

//IContactFactory IContactFactory
type IContactFactory interface {
	New() model.IContact
}
