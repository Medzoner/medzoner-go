package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"github.com/sarulabs/di"
	"os"
)

var TemplaterDefinition = di.Def{
	Name:  "templater",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		pp, _ := os.Getwd()
		return &templater.TemplateHtml{
			RootPath: pp,
		}, nil
	},
}
