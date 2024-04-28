package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/path"
	"github.com/sarulabs/di"
)

// ConfigDependency ConfigDependency
type ConfigDependency struct {
	ConfigInstance config.Config
}

// GetDefinition GetDefinition
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

// InitConfig InitConfig
func (cd *ConfigDependency) InitConfig(rootPath path.RootPath) config.Config {
	c := config.Config{
		RootPath:  string(rootPath),
		DebugMode: false,
		Options:   nil,
		APIPort:   8002,
	}
	c.Init()
	cd.ConfigInstance = c
	return c
}
