package web

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/Medzoner/medzoner-go/pkg/infra/server"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"net/http"
)

//IWeb IWeb
type IWeb interface {
	Start()
}

//Web Web
type Web struct {
	Logger          logger.ILogger
	Router          router.IRouter
	Server          server.IServer
	NotFoundHandler *handler.NotFoundHandler
	IndexHandler    *handler.IndexHandler
	TechnoHandler   *handler.TechnoHandler
	ContactHandler  *handler.ContactHandler
	APIPort         int
}

//Start Start
func (a *Web) Start() {
	a.Router.SetNotFoundHandler(a.NotFoundHandler.Handle)
	a.Router.HandleFunc("/", a.IndexHandler.IndexHandle).Methods("GET")
	a.Router.HandleFunc("/contact", a.ContactHandler.IndexHandle).Methods("POST")
	a.Router.Use(middleware.APIMiddleware{Logger: a.Logger}.Middleware)

	a.Router.PathPrefix("/public").Handler(http.FileServer(http.Dir(".")))
	a.Router.Handle("/")

	_ = a.Logger.Log(fmt.Sprintf("Server up on port '%d'", a.APIPort))
	err := a.Server.ListenAndServe()
	if err != nil {
		_ = a.Logger.Error(fmt.Sprintln(err))
	}
}
