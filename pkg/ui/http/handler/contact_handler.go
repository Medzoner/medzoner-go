package handler

import (
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"time"
)

type ContactHandler struct {
	Template templater.Templater
	CreateContactCommandHandler command.CreateContactCommandHandler
}

type ContactView struct {
	Locale    string
	PageTitle string
	Message   interface{}
	Errors    []validator.FieldError
	TorHost   string
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func (c *ContactHandler) IndexHandle(response http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, "session-name")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	view := ContactView{
		Locale:    "fr",
		PageTitle: "MedZoner.com",
		Message:   session.Values["message"],
	}
	statusCode := http.StatusOK
	if request.Method == "POST" && request.FormValue("Envoyer") == "" {
		createContactCommand := command.CreateContactCommand{
			DateAdd: time.Now(),
			Name:    request.FormValue("name"),
			Email:   request.FormValue("email"),
			Message: request.FormValue("message"),
		}
		v := validator.New()
		err := v.Struct(createContactCommand)
		if err == nil {
			c.CreateContactCommandHandler.Handle(createContactCommand)
			session.Values["message"] = "Votre message a bien été envoyé. Merci!"
			err = session.Save(request, response)
			if err != nil {
				http.Error(response, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(response, request, "/contact", http.StatusSeeOther)
			return
		}
		statusCode = http.StatusBadRequest
		view.Errors = err.(validator.ValidationErrors)
	}
	session.Values["message"] = ""
	err = session.Save(request, response)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	view.TorHost = request.Header.Get("TOR-HOST")
	c.Template.Render("contact", view, response, statusCode)

	_ = request
}
