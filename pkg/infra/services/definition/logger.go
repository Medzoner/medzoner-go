package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/sarulabs/di"
)

var LoggerDefinition = di.Def{
	Name:  "logger",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return logger.ZapLoggerAdapter{
			RootPath: ctn.Get("config").(config.IConfig).GetRootPath() + "/",
			UseSugar: false,
		}.New(), nil
	},
}
