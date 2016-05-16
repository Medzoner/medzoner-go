package definition

import (
	logger "github.com/Medzoner/gologger"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/sarulabs/di"
)

var LoggerDefinition = di.Def{
	Name:  "logger",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return &logger.Logger{
			LogPath: ctn.Get("config").(config.IConfig).GetRootPath() + "/var/log/" + ctn.Get("config").(config.IConfig).GetEnvironment() + ".log",
		}, nil
	},
}
