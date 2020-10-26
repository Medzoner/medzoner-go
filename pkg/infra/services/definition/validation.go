package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/sarulabs/di"
)

//ValidationDefinition ValidationDefinition
var ValidationDefinition = di.Def{
	Name:  "validation",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return validation.ValidatorAdapter{}.New(), nil
	},
}
