package web

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/Medzoner/medzoner-go/pkg/infra/server"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
	"net/http"
	"regexp"
)

// IWeb IWeb
type IWeb interface {
	Start()
}

// Web Web
type Web struct {
	Logger             logger.ILogger
	Router             router.IRouter
	Server             server.IServer
	NotFoundHandler    *handler.NotFoundHandler
	IndexHandler       *handler.IndexHandler
	TechnoHandler      *handler.TechnoHandler
	APIPort            int
	RecaptchaSecretKey string
}

// NewWeb NewWeb
func NewWeb(
	logger logger.ILogger,
	router router.IRouter,
	server server.IServer,
	notFoundHandler *handler.NotFoundHandler,
	indexHandler *handler.IndexHandler,
	technoHandler *handler.TechnoHandler,
	conf config.IConfig,
) *Web {
	return &Web{
		Logger:             logger,
		Router:             router,
		Server:             server,
		NotFoundHandler:    notFoundHandler,
		IndexHandler:       indexHandler,
		TechnoHandler:      technoHandler,
		APIPort:            conf.GetAPIPort(),
		RecaptchaSecretKey: conf.GetRecaptchaSecretKey(),
	}
}

// Start Start
func (a *Web) Start() {
	recaptcha.Init(a.RecaptchaSecretKey)
	a.Router.SetNotFoundHandler(a.NotFoundHandler.Handle)
	a.Router.HandleFunc("/", a.IndexHandler.IndexHandle).Methods("GET", "POST")
	a.Router.Use(middleware.NewAPIMiddleware(a.Logger).Middleware)
	fs := http.FileServer(http.Dir("."))

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	a.Router.PathPrefix("/public/css").Handler(m.Middleware(fs))
	a.Router.PathPrefix("/public/js").Handler(m.Middleware(fs))

	a.Router.PathPrefix("/public").Handler(fs)
	a.Router.Handle("/")

	err := a.Logger.Log(fmt.Sprintf("Server up on port '%d'", a.APIPort))
	if err != nil {
		_ = a.Logger.Error(fmt.Sprintln(err))
	}
	err = a.Server.ListenAndServe()
	if err != nil {
		_ = a.Logger.Error(fmt.Sprintln(err))
	}
}
