package handler

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"github.com/dpapathanasiou/go-recaptcha"
	"log"
	"net/http"
	"time"
)

//IndexHandler IndexHandler
type IndexHandler struct {
	Template                    templater.Templater
	ListTechnoQueryHandler      query.ListTechnoQueryHandler
	RecaptchaSiteKey            string
	CreateContactCommandHandler command.CreateContactCommandHandler
	Session                     session.Sessioner
	Validation                  validation.MzValidator
}

//IndexView IndexView
type IndexView struct {
	Locale    string
	PageTitle string
	TorHost   string
	TechnoView
	Message          interface{}
	Errors           interface{}
	RecaptchaSiteKey string
	PageDescription  string
}

func processRequest(request *http.Request) (result bool) {
	recaptchaResponse, responseFound := request.Form["g-recaptcha-response"]
	if responseFound {
		result, err := recaptcha.Confirm("127.0.0.1", recaptchaResponse[0])
		if err != nil {
			log.Println("recaptcha server error", err)
		}
		return result
	}
	return false
}

//IndexHandle IndexHandle
func (h *IndexHandler) IndexHandle(response http.ResponseWriter, request *http.Request) {
	view := IndexView{
		Locale:    "fr",
		PageTitle: "MedZoner.com",
		TorHost:   request.Header.Get("TOR-HOST"),
		TechnoView: TechnoView{
			Stacks: h.ListTechnoQueryHandler.Handle(query.ListTechnoQuery{Type: "stack"}),
		},
		RecaptchaSiteKey: h.RecaptchaSiteKey,
		PageDescription:  "Mehdi YOUB - Développeur Web Full Stack - NestJS Symfony Golang VueJS",
	}

	newSession, err := h.Session.Init(request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Session = newSession
	h.Session.SetValue("message", "")
	err = h.Session.Save(request, response)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	statusCode := http.StatusOK
	if request.Method == "POST" && request.FormValue("Envoyer") == "" {
		if processRequest(request) {
			createContactCommand := command.CreateContactCommand{
				DateAdd: time.Now(),
				Name:    request.FormValue("name"),
				Email:   request.FormValue("email"),
				Message: request.FormValue("message"),
			}
			v := h.Validation
			err := v.Struct(createContactCommand)
			if err == nil {
				h.CreateContactCommandHandler.Handle(createContactCommand)
				h.Session.SetValue("message", "Votre message a bien été envoyé. Merci!")
				err = h.Session.Save(request, response)
				if err != nil {
					http.Error(response, err.Error(), http.StatusInternalServerError)
					return
				}
				http.Redirect(response, request, "/", http.StatusSeeOther)
				return
			}
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusBadRequest
			fmt.Println("Recaptcha was incorrect; try again.")
			h.Session.SetValue("message", "Recaptcha was incorrect; try again.")
			_ = h.Session.Save(request, response)
			http.Redirect(response, request, "/", http.StatusSeeOther)
			return
		}
	}
	_, err = h.Template.Render("index", view, response, statusCode)
	if err != nil {
		fmt.Println(err)
	}
	_ = request
}
