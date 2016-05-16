package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/security"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/web"
	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
	"net/http"
)

var WebDefinition = di.Def{
	Name:  "app-web",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return &web.Web{
			Logger:         ctn.Get("logger").(logger.ILogger),
			Router:         ctn.Get("router").(*mux.Router),
			Server:         ctn.Get("server").(*http.Server),
			IndexHandler:   ctn.Get("index-handler").(*handler.IndexHandler),
			TechnoHandler:  ctn.Get("techno-handler").(*handler.TechnoHandler),
			ContactHandler: ctn.Get("contact-handler").(*handler.ContactHandler),
			ApiPort:        ctn.Get("config").(config.IConfig).GetApiPort(),
			Security:       ctn.Get("security").(security.IAuth),
		}, nil
	},
}
