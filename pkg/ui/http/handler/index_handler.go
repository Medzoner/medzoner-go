package handler

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/http_utils"
	"net/http"
	"time"

	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/infra/captcha"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
)

// IndexView IndexView
type IndexView struct {
	Locale    string
	PageTitle string
	TorHost   string
	TechnoView
	Errors           interface{}
	RecaptchaSiteKey string
	PageDescription  string
	FormMessage      string
}

// TechnoView TechnoView
type TechnoView struct {
	Locale      string
	PageTitle   string
	Stacks      interface{}
	Experiences interface{}
	Formations  interface{}
	Langs       interface{}
	Others      interface{}
	TorHost     string
}

// IndexHandler IndexHandler
type IndexHandler struct {
	Template                    templater.Templater
	ListTechnoQueryHandler      query.ListTechnoQueryHandler
	RecaptchaSiteKey            string
	CreateContactCommandHandler command.CreateContactCommandHandler
	Session                     session.Sessioner
	Validation                  validation.MzValidator
	Recaptcha                   captcha.Captcher
	Tracer                      tracer.Tracer
	Debug                       bool
}

// NewIndexHandler NewIndexHandler
func NewIndexHandler(
	template templater.Templater,
	listTechnoQueryHandler query.ListTechnoQueryHandler,
	conf config.Config,
	createContactCommandHandler command.CreateContactCommandHandler,
	session session.Sessioner,
	validation validation.MzValidator,
	recaptcha captcha.Captcher,
	tracer tracer.Tracer,
) *IndexHandler {
	return &IndexHandler{
		Template:                    template,
		ListTechnoQueryHandler:      listTechnoQueryHandler,
		RecaptchaSiteKey:            conf.RecaptchaSiteKey,
		CreateContactCommandHandler: createContactCommandHandler,
		Session:                     session,
		Validation:                  validation,
		Recaptcha:                   recaptcha,
		Tracer:                      tracer,
		Debug:                       conf.DebugMode,
	}
}

func (h *IndexHandler) processRequest(request *http.Request) (err error) {
	recaptchaResponse, responseFound := request.Form["g-captcha-response"]
	if responseFound {
		result, err := h.Recaptcha.Confirm(request.RemoteAddr, recaptchaResponse[0])
		if err != nil {
			return fmt.Errorf("captcha server error: %w", err)
		}
		if !result && !h.Debug {
			return fmt.Errorf("captcha was incorrect; try again")
		}
	}
	return nil
}

// IndexHandle IndexHandle
func (h *IndexHandler) IndexHandle(response http.ResponseWriter, request *http.Request) {
	ctx, span := h.Tracer.StartRoot(request.Context(), request, "IndexHandler.IndexHandle")
	defer span.End()

	newSession, err := h.Session.Init(request)
	if err != nil {
		http_utils.ResponseError(response, err, http.StatusInternalServerError, span)
		return
	}

	view, err := h.initView(ctx, request)
	if err != nil {
		http_utils.ResponseError(response, err, http.StatusInternalServerError, span)
		return
	}
	if newSession.GetValue("message") != nil {
		view.FormMessage = newSession.GetValue("message").(string)
	}
	statusCode := http.StatusOK
	if request.Method == "POST" && request.FormValue("submit") == "" {
		if err = h.processRequest(request); err != nil {
			newSession.SetValue("message", "Recaptcha was incorrect; try again.")
			_ = newSession.Save(request, response)
			http.Redirect(response, request, "/#contact", http.StatusSeeOther)
			return
		}

		createContactCommand := command.CreateContactCommand{
			DateAdd: time.Now(),
			Name:    request.FormValue("name"),
			Email:   request.FormValue("email"),
			Message: request.FormValue("message"),
		}

		validationError := h.Validation.Struct(createContactCommand)
		if validationError == nil {
			if err = h.CreateContactCommandHandler.Handle(ctx, createContactCommand); err != nil {
				// newSession.SetValue("message", "Error during send message")
				// if err = newSession.Save(request, response); err != nil {
				//	span.RecordError(err)
				//	http_utils.ResponseError(response, err.Error(), http.StatusInternalServerError)
				//	return
				// }
				// http.Redirect(response, request, "/#contact", http.StatusSeeOther)
				http_utils.ResponseError(response, err, http.StatusInternalServerError, span)
				return
			}
			newSession.SetValue("message", "Your Message has been sent")

			if err = newSession.Save(request, response); err != nil {
				http_utils.ResponseError(response, err, http.StatusInternalServerError, span)
				return
			}
			http.Redirect(response, request, "/#contact", http.StatusSeeOther)
			return
		}
		statusCode = http.StatusBadRequest
	}
	if view.FormMessage != "" {
		response.WriteHeader(statusCode)
		newSession.SetValue("message", "")
		if err = newSession.Save(request, response); err != nil {
			http_utils.ResponseError(response, err, http.StatusInternalServerError, span)
			return
		}
	}
	_, err = h.Template.Render("index", view, response, statusCode)
	if err != nil {
		http_utils.ResponseError(response, err, http.StatusInternalServerError, span)
		return
	}
}

func (h *IndexHandler) initView(ctx context.Context, request *http.Request) (IndexView, error) {
	stacks, err := h.ListTechnoQueryHandler.Handle(ctx, query.ListTechnoQuery{Type: "stack"})
	if err != nil {
		return IndexView{}, fmt.Errorf("error during fetch stack: %w", err)
	}

	return IndexView{
		Locale:    "fr",
		PageTitle: "MedZoner.com",
		TorHost:   request.Header.Get("TOR-HOST"),
		TechnoView: TechnoView{
			Stacks: stacks,
		},
		RecaptchaSiteKey: h.RecaptchaSiteKey,
		PageDescription:  "Mehdi YOUB - Développeur Web Full Stack - NestJS Symfony Golang VueJS",
		FormMessage:      "",
	}, nil
}
