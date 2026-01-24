//go:build wireinject
// +build wireinject

package dependency

import (
	"github.com/Medzoner/medzoner-go/internal/application/command"
	event2 "github.com/Medzoner/medzoner-go/internal/application/event"
	"github.com/Medzoner/medzoner-go/internal/application/query"
	"github.com/Medzoner/medzoner-go/internal/application/service/mailer"
	repository2 "github.com/Medzoner/medzoner-go/internal/domain/repository"
	handler2 "github.com/Medzoner/medzoner-go/internal/ui/http/handler"
	"github.com/Medzoner/medzoner-go/internal/ui/http/templater"
	"github.com/Medzoner/medzoner-go/pkg/infra/captcha"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/notification"
	"github.com/Medzoner/medzoner-go/pkg/infra/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	mockBase "github.com/Medzoner/medzoner-go/test"

	"github.com/Medzoner/gomedz/pkg/http"
	srv "github.com/Medzoner/gomedz/pkg/http/server"
	ginadpt "github.com/Medzoner/gomedz/pkg/http/adapter/gin"

	"github.com/google/wire"
	"github.com/Medzoner/gomedz/pkg/http/probes"
	"github.com/Medzoner/gomedz/pkg/observability"
	"github.com/Medzoner/gomedz/pkg/logger"
	"context"
	"github.com/Medzoner/medzoner-go/test/mocks"
	"github.com/Medzoner/medzoner-go/internal/config"
)

func controllers(p *probes.Handler, a handler2.IndexHandler) []http.Controller {
	return []http.Controller{
		p,
		a,
	}
}

func closers(tl observability.Telemetry) []srv.Closer {
	return []srv.Closer{
		tl,
	}
}

func pingers() probes.Pingers {
	return []probes.Probes{}
}

var (
	CommonWiring = wire.NewSet(
		config.NewConfig2,
		wire.FieldsOf(
			new(*config.Config2),
			"Obs",
			"Engine",
			"Logger",
			"Auth",
			"Server",
		),

		pingers,
		probes.New,
	)
	ServerWiring = wire.NewSet(
		ginadpt.New,
		wire.Bind(new(srv.Enginer), new(*ginadpt.Engine)),
		controllers,
		closers,

		srv.NewServer,
	)
	ObsWiring = wire.NewSet(
		logger.NewLogger,
		observability.NewTelemetry,
	)
	UsecaseWiring = wire.NewSet(
		event2.NewContactCreatedEventHandler,
		command.NewCreateContactCommandHandler,
		query.NewListTechnoQueryHandler,

		wire.Bind(new(event2.IEventHandler), new(*event2.ContactCreatedEventHandler)),
	)
	HandlerWiring = wire.NewSet(
		handler2.NewIndexHandler,
		handler2.NewNotFoundHandler,
	)

	InfraWiring = wire.NewSet(
		config.NewConfig,
		templater.NewTemplateHTML,
		validation.NewValidatorAdapter,
		captcha.NewRecaptchaAdapter,
		wire.Bind(new(templater.Templater), new(*templater.TemplateHTML)),
		wire.Bind(new(validation.MzValidator), new(*validation.ValidatorAdapter)),
		wire.Bind(new(captcha.Captcher), new(*captcha.RecaptchaAdapter)),
	)
	DbWiring = wire.NewSet(
		database.NewDbSQLInstance,

		wire.Bind(new(database.DbInstantiator), new(*database.DbSQLInstance)),
	)
	MailerWiring = wire.NewSet(
		notification.NewMailerSMTP,
		wire.Bind(new(mailer.Mailer), new(*notification.MailerSMTP)),
	)
	MailerMockWiring = wire.NewSet(
		wire.FieldsOf(
			new(*mockBase.Mocks),
			"Mailer",
		),
		wire.Bind(new(mailer.Mailer), new(*mocks.MockMailer)),
	)
	RepositoryWiring = wire.NewSet(
		repository.NewTechnoJSONRepository,
		repository.NewMysqlContactRepository,

		wire.Bind(new(repository2.TechnoRepository), new(*repository.TechnoJSONRepository)),
		wire.Bind(new(repository2.ContactRepository), new(*repository.MysqlContactRepository)),
	)
	RepositoryMockWiring = wire.NewSet(
		wire.FieldsOf(
			new(*mockBase.Mocks),
			"TechnoRepository",
		),
		wire.Bind(new(repository2.TechnoRepository), new(*mocks.MockTechnoRepository)),
		wire.FieldsOf(
			new(*mockBase.Mocks),
			"ContactRepository",
		),
		wire.Bind(new(repository2.ContactRepository), new(*mocks.MockContactRepository)),
	)
	AppWiring = wire.NewSet(
		event2.NewContactCreatedEventHandler,
		command.NewCreateContactCommandHandler,
		query.NewListTechnoQueryHandler,

		wire.Bind(new(event2.IEventHandler), new(*event2.ContactCreatedEventHandler)),
	)
	UiWiring = wire.NewSet(
		handler2.NewIndexHandler,
		handler2.NewNotFoundHandler,
	)
)

func InitDbMigration() (database.DbMigration, error) {
	panic(wire.Build(database.NewDbMigration, DbWiring, InfraWiring))
}

func InitServerTest(ctx context.Context) (srv.Server, error) {
	panic(wire.Build(
		DbWiring,
		InfraWiring,
		MailerMockWiring,
		UsecaseWiring,
		CommonWiring,
		ObsWiring,
		RepositoryMockWiring,
		HandlerWiring,
		ServerWiring,
	))
}

func InitServer(ctx context.Context) (srv.Server, error) {
	panic(wire.Build(
		DbWiring,
		InfraWiring,
		MailerWiring,
		UsecaseWiring,
		CommonWiring,
		ObsWiring,
		RepositoryWiring,
		HandlerWiring,
		ServerWiring,
	))
}
