package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/sarulabs/di"
)

// TracerDefinition TracerDefinition
var TracerDefinition = di.Def{
	Name:  "tracer",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return tracer.NewHttpTracer(), nil
	},
}
