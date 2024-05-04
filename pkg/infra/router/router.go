package router

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/gorilla/mux"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
	"net/http"
	"regexp"
)

type IRouter interface {
	Handle(path string)
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
	TechnoHandler   *handler.TechnoHandler
}

func NewMuxRouterAdapter(
	notFoundHandler *handler.NotFoundHandler,
	indexHandler *handler.IndexHandler,
	technoHandler *handler.TechnoHandler,
) *MuxRouterAdapter {
	rt := &MuxRouterAdapter{
		MuxRouter:       mux.NewRouter(),
		NotFoundHandler: notFoundHandler,
		IndexHandler:    indexHandler,
		TechnoHandler:   technoHandler,
	}
	InitRoutes(rt, notFoundHandler, indexHandler)
	return rt
}

// Handle Handle
func (a MuxRouterAdapter) Handle(path string) {
	http.Handle(path, a)
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
	a.MuxRouter.Use(mwf[0])
}

// SetNotFoundHandler SetNotFoundHandler
func (a MuxRouterAdapter) SetNotFoundHandler(handler func(http.ResponseWriter, *http.Request)) {
	a.MuxRouter.NotFoundHandler = http.HandlerFunc(handler)
}

// ServeHTTP ServeHTTP
func (a MuxRouterAdapter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	a.MuxRouter.ServeHTTP(writer, request)
}

func InitRoutes(a IRouter, notFoundHandler *handler.NotFoundHandler, indexHandler *handler.IndexHandler) {
	a.SetNotFoundHandler(notFoundHandler.Handle)
	a.HandleFunc("/", indexHandler.IndexHandle).Methods("GET", "POST")
	a.Use(middleware.NewAPIMiddleware().Middleware)
	fs := http.FileServer(http.Dir("."))

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	a.PathPrefix("/public/css").Handler(m.Middleware(fs))
	a.PathPrefix("/public/js").Handler(m.Middleware(fs))

	a.PathPrefix("/public").Handler(fs)
	a.Handle("/")
}
