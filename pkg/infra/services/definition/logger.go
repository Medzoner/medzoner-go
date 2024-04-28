package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/path"
	"github.com/sarulabs/di"
)

// LoggerDefinition LoggerDefinition
var LoggerDefinition = di.Def{
	Name:  "logger",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return logger.NewLoggerAdapter(
			path.RootPath(ctn.Get("config").(config.IConfig).GetRootPath()+"/"),
			false,
		), nil
	},
}
