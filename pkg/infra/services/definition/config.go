package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/sarulabs/di"
)

type ConfigDependency struct {
	ConfigInstance config.Config
}

func (cd *ConfigDependency) GetDefinition() di.Def {
	var ConfigDefinition = di.Def{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			c := cd.ConfigInstance
			return &c, nil
		},
	}
	return ConfigDefinition
}

func (cd *ConfigDependency) InitConfig(rootPath string) config.Config {
	c := config.Config{
		RootPath:  rootPath,
		DebugMode: false,
		Options:   nil,
		ApiPort:   8000,
	}
	c.Init()
	cd.ConfigInstance = c
	return c
}
