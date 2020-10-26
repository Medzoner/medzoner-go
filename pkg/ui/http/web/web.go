package web

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/gorilla/mux"
	"net/http"
)

type IWeb interface {
	Start()
}

type Web struct {
	Logger         logger.ILogger
	Router         *mux.Router
	Server         *http.Server
	IndexHandler   *handler.IndexHandler
	TechnoHandler  *handler.TechnoHandler
	ContactHandler *handler.ContactHandler
	APIPORT        int
}

func (a *Web) Start() {

	a.Router.HandleFunc("/", a.IndexHandler.IndexHandle).Methods("GET")
	a.Router.HandleFunc("/contact", a.ContactHandler.IndexHandle).Methods("GET")
	a.Router.HandleFunc("/contact", a.ContactHandler.IndexHandle).Methods("POST")
	a.Router.HandleFunc("/technos", a.TechnoHandler.IndexHandle).Methods("GET")
	a.Router.Use(middleware.ApiMiddleware{Logger: a.Logger}.Middleware)

	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", a.Router)
	a.Logger.Log(fmt.Sprintf("Server up on port '%d'", a.APIPORT))
	err := a.Server.ListenAndServe()
	if err != nil {
		a.Logger.Error(fmt.Sprintln(err))
	}
}
