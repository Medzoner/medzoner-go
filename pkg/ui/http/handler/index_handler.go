package handler

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"log"
	"net/http"
	"time"

	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/infra/captcha"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	otelTrace "go.opentelemetry.io/otel/trace"
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
	Logger                      logger.ILogger
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
	Logger logger.ILogger,
) *IndexHandler {
	_, err := tracer.Int64Counter("run", metric.WithDescription("The number of times the iteration ran"))
	if err != nil {
		log.Fatal(err)
	}
	return &IndexHandler{
		Template:                    template,
		ListTechnoQueryHandler:      listTechnoQueryHandler,
		RecaptchaSiteKey:            conf.RecaptchaSiteKey,
		CreateContactCommandHandler: createContactCommandHandler,
		Session:                     session,
		Validation:                  validation,
		Recaptcha:                   recaptcha,
		Tracer:                      tracer,
		Logger:                      Logger,
		Debug:                       conf.DebugMode,
	}
}

func (h *IndexHandler) processRequest(request *http.Request) (err error) {
	recaptchaResponse, responseFound := request.Form["g-captcha-response"]
	if responseFound {
		result, err := h.Recaptcha.Confirm("127.0.0.1", recaptchaResponse[0])
		if err != nil {
			log.Println("captcha server error", err)
			return err
		}
		if !result && !h.Debug {
			return fmt.Errorf("captcha was incorrect; try again")
		}
	}
	return nil
}

// IndexHandle IndexHandle
func (h *IndexHandler) IndexHandle(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), 60*time.Second)
	defer cancel()

	ctx, span := h.Tracer.Start(
		ctx,
		"IndexHandler.IndexHandle",
		otelTrace.WithSpanKind(otelTrace.SpanKindServer),
		otelTrace.WithNewRoot(),
		otelTrace.WithAttributes([]attribute.KeyValue{
			attribute.String("host", request.Host),
			attribute.String("path", request.URL.Path),
			attribute.String("method", request.Method),
		}...))
	defer func() {
		span.End()
	}()
	correlationID := middleware.GetCorrelationID(ctx)
	span.SetAttributes(attribute.String("correlation_id", correlationID))

	newSession, err := h.Session.Init(request)
	if err != nil {
		http.Error(response, "internal error", http.StatusInternalServerError)
		span.RecordError(err)
		return
	}

	view := h.initView(ctx, request)
	if newSession.GetValue("message") != nil {
		view.FormMessage = newSession.GetValue("message").(string)
	}
	statusCode := http.StatusOK
	if request.Method == "POST" && request.FormValue("submit") == "" {
		if err = h.processRequest(request); err != nil {
			fmt.Println("Recaptcha was incorrect; try again.")
			newSession.SetValue("message", "Recaptcha was incorrect; try again.")
			_ = newSession.Save(request, response)
			request.Header.Set("X-Correlation-ID", correlationID)
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
		if err := v.Struct(createContactCommand); err == nil {
			err = h.CreateContactCommandHandler.Handle(ctx, createContactCommand)
			if err != nil {
				newSession.SetValue("message", "Error during send message")
				if err = newSession.Save(request, response); err != nil {
					http.Error(response, err.Error(), http.StatusInternalServerError)
					span.RecordError(err)
					return
				}
				request.Header.Set("X-Correlation-ID", correlationID)
				http.Redirect(response, request, "/#contact", http.StatusSeeOther)
				return
			}
			newSession.SetValue("message", "Your Message has been sent")

			if err = newSession.Save(request, response); err != nil {
				http.Error(response, err.Error(), http.StatusInternalServerError)
				span.RecordError(err)
				return
			}
			request.Header.Set("X-Correlation-ID", correlationID)
			http.Redirect(response, request, "/#contact", http.StatusSeeOther)
			return
		}
		statusCode = http.StatusBadRequest
	}
	if view.FormMessage != "" {
		response.WriteHeader(statusCode)
		newSession.SetValue("message", "")
		if err = newSession.Save(request, response); err != nil {
			http.Error(response, "internal error", http.StatusInternalServerError)
			span.RecordError(err)
			return
		}
	}
	_, err = h.Template.Render("index", view, response, statusCode)
	if err != nil {
		http.Error(response, "internal error", http.StatusInternalServerError)
		span.RecordError(err)
		return
	}
}

func (h *IndexHandler) initView(ctx context.Context, request *http.Request) IndexView {
	return IndexView{
		Locale:    "fr",
		PageTitle: "MedZoner.com",
		TorHost:   request.Header.Get("TOR-HOST"),
		TechnoView: TechnoView{
			Stacks: h.ListTechnoQueryHandler.Handle(ctx, query.ListTechnoQuery{Type: "stack"}),
		},
		RecaptchaSiteKey: h.RecaptchaSiteKey,
		PageDescription:  "Mehdi YOUB - Développeur Web Full Stack - NestJS Symfony Golang VueJS",
		FormMessage:      "",
	}
}
