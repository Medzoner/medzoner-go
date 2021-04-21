package router

import (
	"github.com/gorilla/mux"
	"net/http"
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
	MuxRouter *mux.Router
}

//Handle Handle
func (a MuxRouterAdapter) Handle(path string) {
	http.Handle(path, a)
}

//HandleFunc HandleFunc
func (a MuxRouterAdapter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return a.MuxRouter.HandleFunc(path, f)
}

//PathPrefix PathPrefix
func (a MuxRouterAdapter) PathPrefix(tpl string) *mux.Route {
	return a.MuxRouter.PathPrefix(tpl)
}

//Use Use
func (a MuxRouterAdapter) Use(mwf ...mux.MiddlewareFunc) {
	a.MuxRouter.Use(mwf[0])
}

//SetNotFoundHandler SetNotFoundHandler
func (a MuxRouterAdapter) SetNotFoundHandler(handler func(http.ResponseWriter, *http.Request)) {
	a.MuxRouter.NotFoundHandler = http.HandlerFunc(handler)
}

//ServeHTTP ServeHTTP
func (a MuxRouterAdapter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	a.MuxRouter.ServeHTTP(writer, request)
}
