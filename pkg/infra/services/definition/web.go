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
		return &web.Web{
			Logger:             ctn.Get("logger").(logger.ILogger),
			Router:             ctn.Get("router").(router.IRouter),
			Server:             ctn.Get("server").(*http.Server),
			NotFoundHandler:    ctn.Get("notfound-handler").(*handler.NotFoundHandler),
			IndexHandler:       ctn.Get("index-handler").(*handler.IndexHandler),
			TechnoHandler:      ctn.Get("techno-handler").(*handler.TechnoHandler),
			APIPort:            ctn.Get("config").(config.IConfig).GetAPIPort(),
			RecaptchaSecretKey: ctn.Get("config").(config.IConfig).GetRecaptchaSecretKey(),
		}, nil
	},
}
