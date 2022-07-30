package handler

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"github.com/dpapathanasiou/go-recaptcha"
	"log"
	"net/http"
	"time"
)

//ContactHandler ContactHandler
type ContactHandler struct {
	Template                    templater.Templater
	CreateContactCommandHandler command.CreateContactCommandHandler
	Session                     session.Sessioner
	Validation                  validation.MzValidator
}

//ContactView ContactView
type ContactView struct {
	Locale    string
	PageTitle string
	Message   interface{}
	Errors    interface{}
	TorHost   string
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
func (c *ContactHandler) IndexHandle(response http.ResponseWriter, request *http.Request) {
	newSession, err := c.Session.Init(request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Session = newSession
	c.Session.SetValue("message", "")
	err = c.Session.Save(request, response)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	statusCode := http.StatusOK
	if request.Method == "POST" && request.FormValue("Envoyer") == "" {
		_, buttonClicked := request.Form["button"]
		if buttonClicked {
			if processRequest(request) {
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
					http.Redirect(response, request, "/", http.StatusSeeOther)
					return
				}
				statusCode = http.StatusBadRequest
			} else {
				statusCode = http.StatusBadRequest
				fmt.Println("Recaptcha was incorrect; try again.")
			}
		}
	}
	response.WriteHeader(statusCode)
	_ = request
}
