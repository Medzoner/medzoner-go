package handler

import (
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type ContactHandler struct {
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

func (c *ContactHandler) IndexHandle(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view := ContactView{
		Locale:    "fr",
		PageTitle: "MedZoner.com",
		Message:   session.Values["message"],
	}
	statusCode := http.StatusOK
	if r.Method == "POST" && r.FormValue("Envoyer") == "" {
		createContactCommand := command.CreateContactCommand{
			DateAdd: time.Now(),
			Name:    r.FormValue("name"),
			Email:   r.FormValue("email"),
			Message: r.FormValue("message"),
		}
		v := validator.New()
		err := v.Struct(createContactCommand)
		if err == nil {
			c.CreateContactCommandHandler.Handle(createContactCommand)
			session.Values["message"] = "Votre message a bien été envoyé. Merci!"
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/contact", http.StatusSeeOther)
			return
		}
		statusCode = http.StatusBadRequest
		view.Errors = err.(validator.ValidationErrors)
	}
	session.Values["message"] = ""
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := template.New("contact template")
	pp, _ := os.Getwd()
	t = template.Must(t.ParseFiles(pp+"/tmpl/base.html", pp+"/tmpl/footer.html", pp+"/tmpl/header.html", pp+"/tmpl/Contact/contact.html", pp+"/tmpl/Contact/contact_form.html"))
	w.WriteHeader(statusCode)

	view.TorHost = r.Header.Get("TOR-HOST")
	err = t.ExecuteTemplate(w, "layout", view)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
	_ = r
}
