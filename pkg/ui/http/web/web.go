package web

import (
	"fmt"
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

//IWeb IWeb
type IWeb interface {
	Start()
}

//Web Web
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

//Start Start
func (a *Web) Start() {
	recaptcha.Init(a.RecaptchaSecretKey)
	a.Router.SetNotFoundHandler(a.NotFoundHandler.Handle)
	a.Router.HandleFunc("/", a.IndexHandler.IndexHandle).Methods("GET", "POST")
	a.Router.Use(middleware.APIMiddleware{Logger: a.Logger}.Middleware)
	fs := http.FileServer(http.Dir("."))

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	a.Router.PathPrefix("/public/css").Handler(m.Middleware(fs))
	a.Router.PathPrefix("/public/js").Handler(m.Middleware(fs))

	a.Router.PathPrefix("/public").Handler(fs)
	a.Router.Handle("/")

	_ = a.Logger.Log(fmt.Sprintf("Server up on port '%d'", a.APIPort))
	err := a.Server.ListenAndServe()
	if err != nil {
		_ = a.Logger.Error(fmt.Sprintln(err))
	}
}
