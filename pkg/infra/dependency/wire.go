//go:build wireinject
// +build wireinject

package dependency

import (
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	mocks "github.com/Medzoner/medzoner-go/test"
	contactMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/domain/repository"

	domainRepository "github.com/Medzoner/medzoner-go/pkg/domain/repository"

	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/application/service/mailer"

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

	"github.com/google/wire"
)

var (
	InfraWiring = wire.NewSet(
		config.NewConfig,
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
		mailersmtp.NewMailerSMTP,

		wire.Bind(new(config.IConfig), new(*config.Config)),
		wire.Bind(new(logger.ILogger), new(*logger.ZapLoggerAdapter)),
		wire.Bind(new(router.IRouter), new(*router.MuxRouterAdapter)),
		wire.Bind(new(server.IServer), new(*server.Server)),
		wire.Bind(new(templater.Templater), new(*templater.TemplateHTML)),
		wire.Bind(new(session.Sessioner), new(*session.SessionerAdapter)),
		wire.Bind(new(validation.MzValidator), new(*validation.ValidatorAdapter)),
		wire.Bind(new(captcha.Captcher), new(*captcha.RecaptchaAdapter)),
		wire.Bind(new(tracer.Tracer), new(*tracer.HttpTracer)),
		wire.Bind(new(mailer.Mailer), new(*mailersmtp.MailerSMTP)),
	)
	DbWiring = wire.NewSet(
		database.NewDbSQLInstance,

		wire.Bind(new(database.IDbInstance), new(*database.DbSQLInstance)),
	)
	RepositoryWiring = wire.NewSet(
		repository.NewTechnoJSONRepository,
		repository.NewMysqlContactRepository,

		wire.Bind(new(domainRepository.TechnoRepository), new(*repository.TechnoJSONRepository)),
		wire.Bind(new(domainRepository.ContactRepository), new(*repository.MysqlContactRepository)),
	)
	RepositoryMockWiring = wire.NewSet(
		repository.NewTechnoJSONRepository,
		wire.Bind(new(domainRepository.TechnoRepository), new(*repository.TechnoJSONRepository)),
		wire.FieldsOf(
			new(mocks.Mocks),
			"ContactRepository",
		),
		wire.Bind(new(domainRepository.ContactRepository), new(*contactMock.MockContactRepository)),
	)
	AppWiring = wire.NewSet(
		event.NewContactCreatedEventHandler,
		command.NewCreateContactCommandHandler,
		query.NewListTechnoQueryHandler,

		wire.Bind(new(event.IEventHandler), new(*event.ContactCreatedEventHandler)),
	)
	UiWiring = wire.NewSet(
		handler.NewIndexHandler,
		handler.NewTechnoHandler,
		handler.NewNotFoundHandler,
	)
)

func InitDbInstance() (*database.DbSQLInstance, error) {
	panic(wire.Build(DbWiring, InfraWiring))
}

func InitDbMigration() (database.DbMigration, error) {
	panic(wire.Build(database.NewDbMigration, DbWiring, InfraWiring))
}

func InitServer() (*server.Server, error) {
	panic(wire.Build(InfraWiring, DbWiring, RepositoryWiring, AppWiring, UiWiring))
}

func InitServerTest(mocks mocks.Mocks) (*server.Server, error) {
	panic(wire.Build(InfraWiring, RepositoryMockWiring, AppWiring, UiWiring))
}
