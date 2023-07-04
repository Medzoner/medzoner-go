package definition

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/sarulabs/di"
	"net/http"
)

// ServerDefinition ServerDefinition
var ServerDefinition = di.Def{
	Name:  "server",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		s := http.Server{
			Addr:    fmt.Sprintf(":%d", ctn.Get("config").(config.IConfig).GetAPIPort()),
			Handler: ctn.Get("router").(router.IRouter),
		}
		return &s, nil
	},
}
