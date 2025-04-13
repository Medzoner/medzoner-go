package router

import (
	"net/http"
	"regexp"

	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/gorilla/mux"
	minify "github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
)

type IRouter interface {
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route
	PathPrefix(tpl string) *mux.Route
	Use(mwf ...mux.MiddlewareFunc)
	SetNotFoundHandler(handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type MuxRouterAdapter struct {
	MuxRouter       *mux.Router
	NotFoundHandler *handler.NotFoundHandler
	IndexHandler    *handler.IndexHandler
	Middlewares     middleware.APIMiddleware
}

func NewMuxRouterAdapter(
	notFoundHandler *handler.NotFoundHandler,
	indexHandler *handler.IndexHandler,
	middlewares middleware.APIMiddleware,
) *MuxRouterAdapter {
	rt := &MuxRouterAdapter{
		MuxRouter:       mux.NewRouter(),
		NotFoundHandler: notFoundHandler,
		IndexHandler:    indexHandler,
		Middlewares:     middlewares,
	}
	InitRoutes(rt)

	return rt
}

func InitRoutes(a *MuxRouterAdapter) {
	a.SetNotFoundHandler(a.NotFoundHandler.Handle)
	a.HandleFunc("/", a.IndexHandler.IndexHandle).Methods("GET", "POST")
	a.Use(a.Middlewares.CorrelationMiddleware, a.Middlewares.LogMiddleware, a.Middlewares.CorsMiddleware)
	fs := http.FileServer(http.Dir("."))

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	a.PathPrefix("/public/css").Handler(m.Middleware(fs))
	a.PathPrefix("/public/js").Handler(m.Middleware(fs))

	a.PathPrefix("/public").Handler(fs)
}

// HandleFunc HandleFunc
func (a MuxRouterAdapter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return a.MuxRouter.HandleFunc(path, f)
}

// PathPrefix PathPrefix
func (a MuxRouterAdapter) PathPrefix(tpl string) *mux.Route {
	return a.MuxRouter.PathPrefix(tpl)
}

// Use Use
func (a MuxRouterAdapter) Use(mwf ...mux.MiddlewareFunc) {
	a.MuxRouter.Use(mwf...)
}

// SetNotFoundHandler SetNotFoundHandler
func (a MuxRouterAdapter) SetNotFoundHandler(handler func(http.ResponseWriter, *http.Request)) {
	a.MuxRouter.NotFoundHandler = http.HandlerFunc(handler)
}

// ServeHTTP ServeHTTP
func (a MuxRouterAdapter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	a.MuxRouter.ServeHTTP(writer, request)
}
