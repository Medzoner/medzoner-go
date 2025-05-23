package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	command2 "github.com/Medzoner/medzoner-go/internal/application/command"
	query2 "github.com/Medzoner/medzoner-go/internal/application/query"
	"github.com/Medzoner/medzoner-go/internal/ui/http/http_utils"
	"github.com/Medzoner/medzoner-go/internal/ui/http/templater"
	"github.com/Medzoner/medzoner-go/pkg/infra/captcha"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/telemetry"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
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
	CreateContactCommandHandler command2.CreateContactCommandHandler
	ListTechnoQueryHandler      query2.ListTechnoQueryHandler
	Template                    templater.Templater
	Validation                  validation.MzValidator
	Recaptcha                   captcha.Captcher
	Tracer                      telemetry.Telemeter
	RecaptchaSiteKey            string
	Debug                       bool
}

// NewIndexHandler NewIndexHandler
func NewIndexHandler(
	template templater.Templater,
	listTechnoQueryHandler query2.ListTechnoQueryHandler,
	conf config.Config,
	createContactCommandHandler command2.CreateContactCommandHandler,
	validation validation.MzValidator,
	recaptcha captcha.Captcher,
	tracer telemetry.Telemeter,
) *IndexHandler {
	return &IndexHandler{
		Template:                    template,
		ListTechnoQueryHandler:      listTechnoQueryHandler,
		RecaptchaSiteKey:            conf.RecaptchaSiteKey,
		CreateContactCommandHandler: createContactCommandHandler,
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

	view, err := h.initView(ctx, request)
	if err != nil {
		http_utils.ResponseError(response, err, http.StatusInternalServerError, span)
		return
	}
	statusCode := http.StatusOK
	if request.Method == "POST" && request.FormValue("submit") == "" {
		if err = h.processRequest(request); err != nil {
			http.Redirect(response, request, "/#contact?msg=\"Recaptcha was incorrect; try again.\"", http.StatusSeeOther)
			return
		}
		createContactCommand := command2.CreateContactCommand{
			DateAdd: time.Now(),
			Name:    request.FormValue("name"),
			Email:   request.FormValue("email"),
			Message: request.FormValue("message"),
		}

		validationError := h.Validation.Struct(createContactCommand)
		if validationError == nil {
			if err = h.CreateContactCommandHandler.Handle(ctx, createContactCommand); err != nil {
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
	}

	if err = h.Template.Render("index", view, response); err != nil {
		http_utils.ResponseError(response, err, http.StatusInternalServerError, span)
		return
	}
}

func (h *IndexHandler) initView(ctx context.Context, request *http.Request) (IndexView, error) {
	stacks, err := h.ListTechnoQueryHandler.Handle(ctx, query2.ListTechnoQuery{Type: "stack"})
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
