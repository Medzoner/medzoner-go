package definition

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/web"
	"github.com/sarulabs/di"
	"net/http"
)

// WebDefinition WebDefinition
var WebDefinition = di.Def{
	Name:  "app-web",
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return web.NewWeb(
			ctn.Get("logger").(logger.ILogger),
			ctn.Get("router").(router.IRouter),
			ctn.Get("server").(*http.Server),
			ctn.Get("notfound-handler").(*handler.NotFoundHandler),
			ctn.Get("index-handler").(*handler.IndexHandler),
			ctn.Get("techno-handler").(*handler.TechnoHandler),
			ctn.Get("config").(config.IConfig).GetAPIPort(),
			ctn.Get("config").(config.IConfig).GetRecaptchaSecretKey(),
		), nil
	},
}
