// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependency

import (
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/application/service/mailer"
	repository2 "github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/captcha"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/mailersmtp"
	"github.com/Medzoner/medzoner-go/pkg/infra/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/Medzoner/medzoner-go/pkg/infra/server"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"github.com/Medzoner/medzoner-go/test"
	"github.com/Medzoner/medzoner-go/test/mocks/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/service/mailer"
	"github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitDbMigration() (database.DbMigration, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return database.DbMigration{}, err
	}
	dbSQLInstance := database.NewDbSQLInstance(configConfig)
	dbMigration := database.NewDbMigration(dbSQLInstance, configConfig)
	return dbMigration, nil
}

func InitServer() (*server.Server, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	templateHTML := templater.NewTemplateHTML(configConfig)
	httpTracer, err := tracer.NewHttpTracer(configConfig)
	if err != nil {
		return nil, err
	}
	notFoundHandler := handler.NewNotFoundHandler(templateHTML, httpTracer)
	useSugar := logger.NewUseSugar()
	zapLoggerAdapter := logger.NewLoggerAdapter(useSugar)
	technoJSONRepository := repository.NewTechnoJSONRepository(zapLoggerAdapter, configConfig)
	listTechnoQueryHandler := query.NewListTechnoQueryHandler(technoJSONRepository, httpTracer)
	dbSQLInstance := database.NewDbSQLInstance(configConfig)
	mysqlContactRepository := repository.NewMysqlContactRepository(dbSQLInstance, zapLoggerAdapter, httpTracer)
	mailerSMTP := mailersmtp.NewMailerSMTP(configConfig, httpTracer)
	contactCreatedEventHandler := event.NewContactCreatedEventHandler(mailerSMTP, zapLoggerAdapter, httpTracer)
	createContactCommandHandler := command.NewCreateContactCommandHandler(mysqlContactRepository, contactCreatedEventHandler, zapLoggerAdapter, httpTracer)
	sessionKey := session.NewSessionKey()
	sessionerAdapter := session.NewSessionerAdapter(sessionKey)
	validatorAdapter := validation.NewValidatorAdapter()
	recaptchaAdapter := captcha.NewRecaptchaAdapter()
	indexHandler := handler.NewIndexHandler(templateHTML, listTechnoQueryHandler, configConfig, createContactCommandHandler, sessionerAdapter, validatorAdapter, recaptchaAdapter, httpTracer, zapLoggerAdapter)
	muxRouterAdapter := router.NewMuxRouterAdapter(notFoundHandler, indexHandler)
	serverServer := server.NewServer(configConfig, muxRouterAdapter, zapLoggerAdapter, httpTracer)
	return serverServer, nil
}

func InitServerTest(mocks2 *mocks.Mocks) (*server.Server, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	templateHTML := templater.NewTemplateHTML(configConfig)
	mockTracer := mocks2.HttpTracer
	notFoundHandler := handler.NewNotFoundHandler(templateHTML, mockTracer)
	mockTechnoRepository := mocks2.TechnoRepository
	listTechnoQueryHandler := query.NewListTechnoQueryHandler(mockTechnoRepository, mockTracer)
	mockContactRepository := mocks2.ContactRepository
	mockMailer := mocks2.Mailer
	useSugar := logger.NewUseSugar()
	zapLoggerAdapter := logger.NewLoggerAdapter(useSugar)
	contactCreatedEventHandler := event.NewContactCreatedEventHandler(mockMailer, zapLoggerAdapter, mockTracer)
	createContactCommandHandler := command.NewCreateContactCommandHandler(mockContactRepository, contactCreatedEventHandler, zapLoggerAdapter, mockTracer)
	sessionKey := session.NewSessionKey()
	sessionerAdapter := session.NewSessionerAdapter(sessionKey)
	validatorAdapter := validation.NewValidatorAdapter()
	recaptchaAdapter := captcha.NewRecaptchaAdapter()
	indexHandler := handler.NewIndexHandler(templateHTML, listTechnoQueryHandler, configConfig, createContactCommandHandler, sessionerAdapter, validatorAdapter, recaptchaAdapter, mockTracer, zapLoggerAdapter)
	muxRouterAdapter := router.NewMuxRouterAdapter(notFoundHandler, indexHandler)
	serverServer := server.NewServer(configConfig, muxRouterAdapter, zapLoggerAdapter, mockTracer)
	return serverServer, nil
}

// wire.go:

var (
	InfraWiring      = wire.NewSet(config.NewConfig, logger.NewUseSugar, logger.NewLoggerAdapter, router.NewMuxRouterAdapter, server.NewServer, templater.NewTemplateHTML, session.NewSessionKey, session.NewSessionerAdapter, validation.NewValidatorAdapter, captcha.NewRecaptchaAdapter, wire.Bind(new(logger.ILogger), new(*logger.ZapLoggerAdapter)), wire.Bind(new(router.IRouter), new(*router.MuxRouterAdapter)), wire.Bind(new(server.IServer), new(*server.Server)), wire.Bind(new(templater.Templater), new(*templater.TemplateHTML)), wire.Bind(new(session.Sessioner), new(*session.SessionerAdapter)), wire.Bind(new(validation.MzValidator), new(*validation.ValidatorAdapter)), wire.Bind(new(captcha.Captcher), new(*captcha.RecaptchaAdapter)))
	DbWiring         = wire.NewSet(database.NewDbSQLInstance, wire.Bind(new(database.DbInstantiator), new(*database.DbSQLInstance)))
	TracerWiring     = wire.NewSet(tracer.NewHttpTracer, wire.Bind(new(tracer.Tracer), new(*tracer.HttpTracer)))
	TracerMockWiring = wire.NewSet(wire.FieldsOf(
		new(*mocks.Mocks),
		"HttpTracer",
	), wire.Bind(new(tracer.Tracer), new(*tracerMock.MockTracer)),
	)
	MailerWiring     = wire.NewSet(mailersmtp.NewMailerSMTP, wire.Bind(new(mailer.Mailer), new(*mailersmtp.MailerSMTP)))
	MailerMockWiring = wire.NewSet(wire.FieldsOf(
		new(*mocks.Mocks),
		"Mailer",
	), wire.Bind(new(mailer.Mailer), new(*mailerMock.MockMailer)),
	)
	RepositoryWiring     = wire.NewSet(repository.NewTechnoJSONRepository, repository.NewMysqlContactRepository, wire.Bind(new(repository2.TechnoRepository), new(*repository.TechnoJSONRepository)), wire.Bind(new(repository2.ContactRepository), new(*repository.MysqlContactRepository)))
	RepositoryMockWiring = wire.NewSet(wire.FieldsOf(
		new(*mocks.Mocks),
		"TechnoRepository",
	), wire.Bind(new(repository2.TechnoRepository), new(*domainRepositoryMock.MockTechnoRepository)), wire.FieldsOf(
		new(*mocks.Mocks),
		"ContactRepository",
	), wire.Bind(new(repository2.ContactRepository), new(*domainRepositoryMock.MockContactRepository)),
	)
	AppWiring = wire.NewSet(event.NewContactCreatedEventHandler, command.NewCreateContactCommandHandler, query.NewListTechnoQueryHandler, wire.Bind(new(event.IEventHandler), new(*event.ContactCreatedEventHandler)))
	UiWiring  = wire.NewSet(handler.NewIndexHandler, handler.NewNotFoundHandler)
)
