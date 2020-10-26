package definition

import (
	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

//RouterDefinition RouterDefinition
var RouterDefinition = di.Def{
	Name:  "router",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		r := mux.Router{}
		return &r, nil
	},
}
