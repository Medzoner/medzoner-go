package handler

import (
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"net/http"
	"time"
)

type ContactHandler struct {
	Template                    templater.Templater
	CreateContactCommandHandler command.CreateContactCommandHandler
	Session                     session.Sessioner
	Validation                  validation.MzValidator
}

type ContactView struct {
	Locale    string
	PageTitle string
	Message   interface{}
	Errors    interface{}
	TorHost   string
}

func (c *ContactHandler) IndexHandle(response http.ResponseWriter, request *http.Request) {
	c.Session = c.Session.Init(request)
	view := ContactView{
		Locale:    "fr",
		PageTitle: "MedZoner.com",
		Message:   c.Session.GetValue("message"),
	}
	statusCode := http.StatusOK
	if request.Method == "POST" && request.FormValue("Envoyer") == "" {
		createContactCommand := command.CreateContactCommand{
			DateAdd: time.Now(),
			Name:    request.FormValue("name"),
			Email:   request.FormValue("email"),
			Message: request.FormValue("message"),
		}
		v := c.Validation
		err := v.Struct(createContactCommand)
		if err == nil {
			c.CreateContactCommandHandler.Handle(createContactCommand)
			c.Session.SetValue("message", "Votre message a bien été envoyé. Merci!")
			err = c.Session.Save(request, response)
			if err != nil {
				http.Error(response, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(response, request, "/contact", http.StatusSeeOther)
			return
		}
		statusCode = http.StatusBadRequest
		vErrors := v.GetErrors()
		view.Errors = vErrors
	}
	c.Session.SetValue("message", "")
	err := c.Session.Save(request, response)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.WriteHeader(statusCode)

	view.TorHost = request.Header.Get("TOR-HOST")
	_, err = c.Template.Render("contact", view, response, statusCode)
	if err != nil {
		panic(err)
	}

	_ = request
}
