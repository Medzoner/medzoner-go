package handler

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	otelTrace "go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
	"time"

	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/infra/captcha"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
)

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
	Logger                      logger.ILogger
}

var runCount metric.Int64Counter
var err error

// NewIndexHandler NewIndexHandler
func NewIndexHandler(
	template templater.Templater,
	listTechnoQueryHandler query.ListTechnoQueryHandler,
	conf config.IConfig,
	createContactCommandHandler command.CreateContactCommandHandler,
	session session.Sessioner,
	validation validation.MzValidator,
	recaptcha captcha.Captcher,
	tracer tracer.Tracer,
	Logger logger.ILogger,
) *IndexHandler {
	runCount, err = tracer.Int64Counter("run", metric.WithDescription("The number of times the iteration ran"))
	if err != nil {
		log.Fatal(err)
	}
	return &IndexHandler{
		Template:                    template,
		ListTechnoQueryHandler:      listTechnoQueryHandler,
		RecaptchaSiteKey:            conf.GetRecaptchaSiteKey(),
		CreateContactCommandHandler: createContactCommandHandler,
		Session:                     session,
		Validation:                  validation,
		Recaptcha:                   recaptcha,
		Tracer:                      tracer,
		Logger:                      Logger,
	}
}

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

func (h *IndexHandler) processRequest(request *http.Request) (err error) {
	recaptchaResponse, responseFound := request.Form["g-captcha-response"]
	if responseFound {
		result, err := h.Recaptcha.Confirm("127.0.0.1", recaptchaResponse[0])
		if err != nil {
			log.Println("captcha server error", err)
			return err
		}
		if !result {
			return fmt.Errorf("captcha was incorrect; try again")
		}
	}
	return nil
}

// IndexHandle IndexHandle
func (h *IndexHandler) IndexHandle(response http.ResponseWriter, request *http.Request) {
	contextIndex, cancel := context.WithTimeout(request.Context(), 60*time.Second)
	defer cancel()

	// Attributes represent additional key-value descriptors that can be bound
	// to a metric observer or recorder.
	commonAttrs := []attribute.KeyValue{
		attribute.String("attrA", "chocolate"),
		attribute.String("attrB", "raspberry"),
		attribute.String("attrC", "vanilla"),
	}

	// Work begins
	ctx, span := h.Tracer.Start(
		contextIndex,
		"CollectorExporter-Example",
		otelTrace.WithAttributes(commonAttrs...))
	defer func() {
		span.End()
	}()
	for i := 0; i < 10; i++ {
		_, iSpan := h.Tracer.Start(ctx, fmt.Sprintf("Sample-%d", i))
		runCount.Add(ctx, 1, metric.WithAttributes(commonAttrs...))
		h.Logger.Log(fmt.Sprintf("Doing really hard work (%d / 10)\n", i+1))

		//<-time.After(time.Second)
		iSpan.End()
	}

	log.Printf("Done!")

	h.Tracer.WriteLog(contextIndex, "IndexHandle start")
	newSession, err := h.Session.Init(request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		panic(err.Error())
	}
	h.Session = newSession
	view := IndexView{
		Locale:    "fr",
		PageTitle: "MedZoner.com",
		TorHost:   request.Header.Get("TOR-HOST"),
		TechnoView: TechnoView{
			Stacks: h.ListTechnoQueryHandler.Handle(query.ListTechnoQuery{Type: "stack"}),
		},
		RecaptchaSiteKey: h.RecaptchaSiteKey,
		PageDescription:  "Mehdi YOUB - Développeur Web Full Stack - NestJS Symfony Golang VueJS",
		FormMessage:      "",
	}
	if h.Session.GetValue("message") != nil {
		view.FormMessage = h.Session.GetValue("message").(string)
	}
	statusCode := http.StatusOK
	if request.Method == "POST" && request.FormValue("submit") == "" {
		err = h.processRequest(request)
		if err != nil {
			fmt.Println("Recaptcha was incorrect; try again.")
			h.Session.SetValue("message", "Recaptcha was incorrect; try again.")
			_ = h.Session.Save(request, response)
			http.Redirect(response, request, "/#contact", http.StatusSeeOther)
			return
		}

		createContactCommand := command.CreateContactCommand{
			DateAdd: time.Now(),
			Name:    request.FormValue("name"),
			Email:   request.FormValue("email"),
			Message: request.FormValue("message"),
		}
		v := h.Validation
		err := v.Struct(createContactCommand)
		if err == nil {
			h.CreateContactCommandHandler.Handle(contextIndex, createContactCommand)
			h.Session.SetValue("message", "Your Message has been sent")
			err = h.Session.Save(request, response)
			if err != nil {
				http.Error(response, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(response, request, "/#contact", http.StatusSeeOther)
			return
		}
		statusCode = http.StatusBadRequest
	}
	if view.FormMessage != "" {
		response.WriteHeader(statusCode)
		h.Session.SetValue("message", "")
		err = h.Session.Save(request, response)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	_, err = h.Template.Render("index", view, response, statusCode)
	if err != nil {
		panic(err.Error())
	}

	view.FormMessage = ""
	_ = request
	h.Tracer.WriteLog(contextIndex, "IndexHandle end")
}
