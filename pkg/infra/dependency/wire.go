//go:build wireinject
// +build wireinject

package dependency

import (
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/application/event"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/application/service/mailer"
	domainRepository "github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/captcha"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/infra/notification"
	"github.com/Medzoner/medzoner-go/pkg/infra/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/Medzoner/medzoner-go/pkg/infra/server"
	"github.com/Medzoner/medzoner-go/pkg/infra/telemetry"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	mocks "github.com/Medzoner/medzoner-go/test"
	domainRepositoryMock "github.com/Medzoner/medzoner-go/test/mocks"
	mailerMock "github.com/Medzoner/medzoner-go/test/mocks"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks"

	"github.com/google/wire"
)

var (
	InfraWiring = wire.NewSet(
		config.NewConfig,
		router.NewMuxRouterAdapter,
		server.NewServer,
		templater.NewTemplateHTML,
		validation.NewValidatorAdapter,
		captcha.NewRecaptchaAdapter,
		middleware.NewAPIMiddleware,

		wire.Bind(new(router.IRouter), new(*router.MuxRouterAdapter)),
		wire.Bind(new(server.IServer), new(*server.Server)),
		wire.Bind(new(templater.Templater), new(*templater.TemplateHTML)),
		wire.Bind(new(validation.MzValidator), new(*validation.ValidatorAdapter)),
		wire.Bind(new(captcha.Captcher), new(*captcha.RecaptchaAdapter)),
	)
	DbWiring = wire.NewSet(
		database.NewDbSQLInstance,

		wire.Bind(new(database.DbInstantiator), new(*database.DbSQLInstance)),
	)
	TracerWiring = wire.NewSet(
		telemetry.NewHttpTelemetry,

		wire.Bind(new(telemetry.Telemeter), new(*telemetry.HttpTelemetry)),
	)
	TracerMockWiring = wire.NewSet(
		wire.FieldsOf(
			new(*mocks.Mocks),
			"HttpTelemetry",
		),
		wire.Bind(new(telemetry.Telemeter), new(*tracerMock.MockTelemeter)),
	)
	MailerWiring = wire.NewSet(
		notification.NewMailerSMTP,
		wire.Bind(new(mailer.Mailer), new(*notification.MailerSMTP)),
	)
	MailerMockWiring = wire.NewSet(
		wire.FieldsOf(
			new(*mocks.Mocks),
			"Mailer",
		),
		wire.Bind(new(mailer.Mailer), new(*mailerMock.MockMailer)),
	)
	RepositoryWiring = wire.NewSet(
		repository.NewTechnoJSONRepository,
		repository.NewMysqlContactRepository,

		wire.Bind(new(domainRepository.TechnoRepository), new(*repository.TechnoJSONRepository)),
		wire.Bind(new(domainRepository.ContactRepository), new(*repository.MysqlContactRepository)),
	)
	RepositoryMockWiring = wire.NewSet(
		wire.FieldsOf(
			new(*mocks.Mocks),
			"TechnoRepository",
		),
		wire.Bind(new(domainRepository.TechnoRepository), new(*domainRepositoryMock.MockTechnoRepository)),
		wire.FieldsOf(
			new(*mocks.Mocks),
			"ContactRepository",
		),
		wire.Bind(new(domainRepository.ContactRepository), new(*domainRepositoryMock.MockContactRepository)),
	)
	AppWiring = wire.NewSet(
		event.NewContactCreatedEventHandler,
		command.NewCreateContactCommandHandler,
		query.NewListTechnoQueryHandler,

		wire.Bind(new(event.IEventHandler), new(*event.ContactCreatedEventHandler)),
	)
	UiWiring = wire.NewSet(
		handler.NewIndexHandler,
		handler.NewNotFoundHandler,
	)
)

func InitDbMigration() (database.DbMigration, error) {
	panic(wire.Build(database.NewDbMigration, DbWiring, InfraWiring))
}

func InitServer() (*server.Server, error) {
	panic(wire.Build(InfraWiring, TracerWiring, MailerWiring, DbWiring, RepositoryWiring, AppWiring, UiWiring))
}

func InitServerTest(mocks *mocks.Mocks) (*server.Server, error) {
	panic(wire.Build(InfraWiring, TracerMockWiring, MailerMockWiring, RepositoryMockWiring, AppWiring, UiWiring))
}
