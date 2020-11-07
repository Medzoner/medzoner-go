package web

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/gorilla/mux"
	"net/http"
)

//IWeb IWeb
type IWeb interface {
	Start()
}

//Web Web
type Web struct {
	Logger          logger.ILogger
	Router          *mux.Router
	Server          *http.Server
	NotFoundHandler *handler.NotFoundHandler
	IndexHandler    *handler.IndexHandler
	TechnoHandler   *handler.TechnoHandler
	ContactHandler  *handler.ContactHandler
	APIPORT         int
}

//Start Start
func (a *Web) Start() {

	a.Router.NotFoundHandler = http.HandlerFunc(a.NotFoundHandler.Handle)
	a.Router.HandleFunc("/", a.IndexHandler.IndexHandle).Methods("GET")
	a.Router.HandleFunc("/contact", a.ContactHandler.IndexHandle).Methods("GET")
	a.Router.HandleFunc("/contact", a.ContactHandler.IndexHandle).Methods("POST")
	a.Router.HandleFunc("/technos", a.TechnoHandler.IndexHandle).Methods("GET")
	a.Router.Use(middleware.APIMiddleware{Logger: a.Logger}.Middleware)

	a.Router.PathPrefix("/public").Handler(http.FileServer(http.Dir(".")))
	http.Handle("/", a.Router)

	a.Logger.Log(fmt.Sprintf("Server up on port '%d'", a.APIPORT))
	err := a.Server.ListenAndServe()
	if err != nil {
		a.Logger.Error(fmt.Sprintln(err))
	}
}
