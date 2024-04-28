package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/path"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"github.com/sarulabs/di"
)

// TemplaterDefinition TemplaterDefinition
var TemplaterDefinition = di.Def{
	Name:  "templater",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return templater.NewTemplateHTML(
			path.RootPath(ctn.Get("config").(config.IConfig).GetRootPath() + "/"),
		), nil
	},
}
