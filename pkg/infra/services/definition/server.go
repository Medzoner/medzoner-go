package definition

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
	"net/http"
)

var ServerDefinition = di.Def{
	Name:  "server",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		s := http.Server{
			Addr:    fmt.Sprintf(":%d", ctn.Get("config").(config.IConfig).GetApiPort()),
			Handler: ctn.Get("router").(*mux.Router),
		}
		return &s, nil
	},
}
