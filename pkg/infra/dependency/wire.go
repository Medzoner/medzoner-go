//go:build wireinject
// +build wireinject

package dependency

import (
	"github.com/Medzoner/medzoner-go/pkg/app"
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/application/service/mailer"
	domainRepository "github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/captcha"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/mailersmtp"
	"github.com/Medzoner/medzoner-go/pkg/infra/path"
	"github.com/Medzoner/medzoner-go/pkg/infra/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/Medzoner/medzoner-go/pkg/infra/server"
	"github.com/Medzoner/medzoner-go/pkg/infra/session"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/web"

	"github.com/google/wire"
)

var (
	InfraWiring = wire.NewSet(
		config.NewConfig,
		path.NewRootPath,
		logger.NewUseSugar,
		logger.NewLoggerAdapter,
		router.NewMuxRouterAdapter,
		server.NewServer,
		templater.NewTemplateHTML,
		session.NewSessionKey,
		session.NewSessionerAdapter,
		validation.NewValidatorAdapter,
		captcha.NewRecaptchaAdapter,
		tracer.NewHttpTracer,
		database.NewDbSQLInstance,
		database.NewDbMigration,
		mailersmtp.NewMailerSMTP,
		repository.NewTechnoJSONRepository,
		repository.NewMysqlContactRepository,
		wire.Bind(new(config.IConfig), new(*config.Config)),
		wire.Bind(new(logger.ILogger), new(*logger.ZapLoggerAdapter)),
		wire.Bind(new(router.IRouter), new(*router.MuxRouterAdapter)),
		wire.Bind(new(server.IServer), new(*server.Server)),
		wire.Bind(new(templater.Templater), new(*templater.TemplateHTML)),
		wire.Bind(new(session.Sessioner), new(*session.SessionerAdapter)),
		wire.Bind(new(validation.MzValidator), new(*validation.ValidatorAdapter)),
		wire.Bind(new(captcha.Captcher), new(*captcha.RecaptchaAdapter)),
		wire.Bind(new(tracer.Tracer), new(*tracer.HttpTracer)),
		wire.Bind(new(database.IDbInstance), new(*database.DbSQLInstance)),
		wire.Bind(new(mailer.Mailer), new(*mailersmtp.MailerSMTP)),
		wire.Bind(new(domainRepository.TechnoRepository), new(*repository.TechnoJSONRepository)),
		wire.Bind(new(domainRepository.ContactRepository), new(*repository.MysqlContactRepository)),
	)
	AppWiring = wire.NewSet(
		event.NewContactCreatedEventHandler,
		wire.Bind(new(event.IEventHandler), new(*event.ContactCreatedEventHandler)),
		command.NewCreateContactCommandHandler,
		query.NewListTechnoQueryHandler,
	)
	UiWiring = wire.NewSet(
		handler.NewIndexHandler,
		handler.NewTechnoHandler,
		handler.NewNotFoundHandler,
	)
	CoreWiring = wire.NewSet(
		web.NewWeb,
		wire.Bind(new(web.IWeb), new(*web.Web)),
	)
)

func InitApp() *app.App {
	panic(wire.Build(app.NewApp, InfraWiring, AppWiring, UiWiring, CoreWiring))
}

//var (
//ConfigWiring     = wire.NewSet(config.NewConfig)
//RootPathWiring   = wire.NewSet(path.NewRootPath)
//UseSugarWiring   = wire.NewSet(logger.NewUseSugar)
//LoggerWiring     = wire.NewSet(logger.NewLoggerAdapter, wire.Bind(new(logger.ILogger), new(*logger.ZapLoggerAdapter)))
//RouterWiring     = wire.NewSet(router.NewMuxRouterAdapter, wire.Bind(new(router.IRouter), new(*router.MuxRouterAdapter)))
//ServerWiring     = wire.NewSet(server.NewServer, wire.Bind(new(server.IServer), new(*server.Server)))
//TemplaterWiring  = wire.NewSet(templater.NewTemplateHTML, wire.Bind(new(templater.Templater), new(*templater.TemplateHTML)))
//SessionKeyWiring = wire.NewSet(session.NewSessionKey)
//SessionWiring    = wire.NewSet(session.NewSessionerAdapter, wire.Bind(new(session.Sessioner), new(*session.SessionerAdapter)))
//ValidationWiring = wire.NewSet(validation.NewValidatorAdapter, wire.Bind(new(validation.MzValidator), new(*validation.ValidatorAdapter)))
//RecaptchaWiring  = wire.NewSet(captcha.NewRecaptchaAdapter, wire.Bind(new(captcha.Captcher), new(*captcha.RecaptchaAdapter)))
//TracerWiring     = wire.NewSet(tracer.NewHttpTracer, wire.Bind(new(tracer.Tracer), new(*tracer.HttpTracer)))
//DbInstanceWiring = wire.NewSet(database.NewDbSQLInstance, wire.Bind(new(database.IDbInstance), new(*database.DbSQLInstance)))
//DbMigrateWiring  = wire.NewSet(database.NewDbMigration)
//MailerWiring     = wire.NewSet(mailersmtp.NewMailerSMTP, wire.Bind(new(mailer.Mailer), new(*mailersmtp.MailerSMTP)))

//TechnoRepositoryWiring  = wire.NewSet(repository.NewTechnoJSONRepository, wire.Bind(new(domainRepository.TechnoRepository), new(*repository.TechnoJSONRepository)))
//ContactRepositoryWiring = wire.NewSet(repository.NewMysqlContactRepository, wire.Bind(new(domainRepository.ContactRepository), new(*repository.MysqlContactRepository)))

//ContactCreatedEventHandlerWiring  = wire.NewSet(event.NewContactCreatedEventHandler, wire.Bind(new(event.IEventHandler), new(*event.ContactCreatedEventHandler)))
//ListTechnoQueryHandlerWiring      = wire.NewSet(query.NewListTechnoQueryHandler)
//CreateContactCommandHandlerWiring = wire.NewSet(command.NewCreateContactCommandHandler)

//IndexHandlerWiring    = wire.NewSet(handler.NewIndexHandler)
//NotFoundHandlerWiring = wire.NewSet(handler.NewNotFoundHandler)
//TechnoHandlerWiring   = wire.NewSet(handler.NewTechnoHandler)

//WebWiring = wire.NewSet(web.NewWeb)
//AppWiring = wire.NewSet(app.NewApp)
//)

//func InitRootPath() path.RootPath {
//	panic(wire.Build(InfraWiring))
//}
//func InitConfig() config.IConfig {
//	panic(wire.Build(InfraWiring))
//}
//func InitUseSugar() logger.UseSugar {
//	panic(wire.Build(InfraWiring))
//}
//func InitLogger() logger.ILogger {
//	panic(wire.Build(InfraWiring, InitRootPath, InitUseSugar))
//}
//func InitRouter() router.IRouter {
//	panic(wire.Build(InfraWiring))
//}
//func InitServer() server.IServer {
//	panic(wire.Build(InfraWiring, InitRouter, InitConfig))
//}
//func InitTemplater() templater.Templater {
//	panic(wire.Build(InfraWiring, InitRootPath))
//}
//func InitSessionKey() session.SessionKey {
//	panic(wire.Build(InfraWiring))
//}
//func InitSession() session.Sessioner {
//	panic(wire.Build(InfraWiring, InitSessionKey))
//}
//func InitValidation() validation.MzValidator {
//	panic(wire.Build(InfraWiring))
//}
//func InitRecaptcha() captcha.Captcher {
//	panic(wire.Build(InfraWiring))
//}
//func InitTracer() tracer.Tracer {
//	panic(wire.Build(InfraWiring))
//}
//func InitDbInstance() database.IDbInstance {
//	panic(wire.Build(InfraWiring, InitConfig))
//}
//func InitDbMigration() database.DbMigration {
//	panic(wire.Build(InfraWiring, InitDbInstance, InitRootPath))
//}
//func InitMailer() mailer.Mailer {
//	panic(wire.Build(InfraWiring, InitConfig, InitRootPath))
//}

//func InitTechnoRepository() domainRepository.TechnoRepository {
//	panic(wire.Build(TechnoRepositoryWiring, InfraWiring))
//}
//func InitContactRepository() domainRepository.ContactRepository {
//	panic(wire.Build(ContactRepositoryWiring, InfraWiring))
//}

//func InitContactCreatedEventHandler() event.IEventHandler {
//	panic(wire.Build(ContactCreatedEventHandlerWiring, InfraWiring))
//}
//func InitListTechnoQueryHandler() query.ListTechnoQueryHandler {
//	panic(wire.Build(ListTechnoQueryHandlerWiring, InfraWiring))
//}
//func InitCreateContactCommandHandler() command.CreateContactCommandHandler {
//	panic(wire.Build(CreateContactCommandHandlerWiring, InfraWiring, InitContactCreatedEventHandler))
//}

//func InitIndexHandler() *handler.IndexHandler {
//	panic(wire.Build(
//		IndexHandlerWiring,
//		InfraWiring,
//		InitListTechnoQueryHandler,
//		InitCreateContactCommandHandler,
//	))
//}
//func InitNotFoundHandler() *handler.NotFoundHandler {
//	panic(wire.Build(NotFoundHandlerWiring, InfraWiring))
//}
//func InitTechnoHandler() *handler.TechnoHandler {
//	panic(wire.Build(TechnoHandlerWiring, InfraWiring, InitListTechnoQueryHandler))
//}

//	func IniWeb() web.IWeb {
//		panic(wire.Build(AppWiring, InfraWiring, InitIndexHandler, InitNotFoundHandler, InitTechnoHandler))
//	}
