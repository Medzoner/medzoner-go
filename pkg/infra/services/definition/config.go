package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/sarulabs/di"
	"os"
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

func (cd *ConfigDependency) InitConfig() config.Config {
	rootPath, _ := os.Getwd()
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
