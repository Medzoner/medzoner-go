package server

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
	"net/http"
	"regexp"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
)

// IServer Server Server
type IServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

// Server Server
type Server struct {
	Logger             logger.ILogger
	Router             router.IRouter
	HTTPServer         *http.Server
	APIPort            int
	RecaptchaSecretKey string
	NotFoundHandler    *handler.NotFoundHandler
	IndexHandler       *handler.IndexHandler
	TechnoHandler      *handler.TechnoHandler
}

// NewServer NewServer
func NewServer(
	conf config.IConfig,
	route router.IRouter,
	logger logger.ILogger,
	notFoundHandler *handler.NotFoundHandler,
	indexHandler *handler.IndexHandler,
	technoHandler *handler.TechnoHandler,
) *Server {
	return &Server{
		Logger: logger,
		Router: route,
		HTTPServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", conf.GetAPIPort()),
			Handler: route,
		},
		APIPort:            conf.GetAPIPort(),
		RecaptchaSecretKey: conf.GetRecaptchaSecretKey(),
		NotFoundHandler:    notFoundHandler,
		IndexHandler:       indexHandler,
		TechnoHandler:      technoHandler,
	}
}

func (s Server) ListenAndServe() error {
	recaptcha.Init(s.RecaptchaSecretKey)
	s.Router.SetNotFoundHandler(s.NotFoundHandler.Handle)
	s.Router.HandleFunc("/", s.IndexHandler.IndexHandle).Methods("GET", "POST")
	s.Router.Use(middleware.NewAPIMiddleware(s.Logger).Middleware)
	fs := http.FileServer(http.Dir("."))

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	s.Router.PathPrefix("/public/css").Handler(m.Middleware(fs))
	s.Router.PathPrefix("/public/js").Handler(m.Middleware(fs))

	s.Router.PathPrefix("/public").Handler(fs)
	s.Router.Handle("/")

	err := s.Logger.Log(fmt.Sprintf("Server up on port '%d'", s.APIPort))
	if err != nil {
		_ = s.Logger.Error(fmt.Sprintln(err))
	}
	return s.HTTPServer.ListenAndServe()
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.HTTPServer.Shutdown(ctx)
}
