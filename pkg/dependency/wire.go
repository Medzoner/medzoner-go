//go:build wireinject

package dependency

import (
	"github.com/Medzoner/medzoner-go/internal/application/command"
	event2 "github.com/Medzoner/medzoner-go/internal/application/event"
	"github.com/Medzoner/medzoner-go/internal/application/query"
	"github.com/Medzoner/medzoner-go/internal/application/service/mailer"
	repository2 "github.com/Medzoner/medzoner-go/internal/domain/repository"
	handler2 "github.com/Medzoner/medzoner-go/internal/ui/http/handler"
	"github.com/Medzoner/medzoner-go/internal/ui/http/templater"
	mockBase "github.com/Medzoner/medzoner-go/test"

	"github.com/Medzoner/gomedz/pkg/http"
	ginadpt "github.com/Medzoner/gomedz/pkg/http/adapter/gin"
	srv "github.com/Medzoner/gomedz/pkg/http/server"

	"context"
	"github.com/Medzoner/gomedz/pkg/http/probes"
	"github.com/Medzoner/gomedz/pkg/logger"
	"github.com/Medzoner/gomedz/pkg/observability"
	"github.com/Medzoner/medzoner-go/internal/config"
	"github.com/Medzoner/medzoner-go/pkg/captcha"
	database2 "github.com/Medzoner/medzoner-go/pkg/database"
	"github.com/Medzoner/medzoner-go/pkg/notification"
	repository3 "github.com/Medzoner/medzoner-go/pkg/repository"
	"github.com/Medzoner/medzoner-go/pkg/validation"
	"github.com/Medzoner/medzoner-go/test/mocks"
	"github.com/google/wire"
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
		config.NewConfig,
		wire.FieldsOf(
			new(config.Config),
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
		templater.NewTemplateHTML,
		validation.NewValidatorAdapter,
		captcha.NewRecaptchaAdapter,
		wire.Bind(new(templater.Templater), new(*templater.TemplateHTML)),
		wire.Bind(new(validation.MzValidator), new(*validation.ValidatorAdapter)),
		wire.Bind(new(captcha.Captcher), new(*captcha.RecaptchaAdapter)),
	)
	DbWiring = wire.NewSet(
		database2.NewDbSQLInstance,

		wire.Bind(new(database2.DbInstantiator), new(*database2.DbSQLInstance)),
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
		repository3.NewTechnoJSONRepository,
		repository3.NewMysqlContactRepository,

		wire.Bind(new(repository2.TechnoRepository), new(*repository3.TechnoJSONRepository)),
		wire.Bind(new(repository2.ContactRepository), new(*repository3.MysqlContactRepository)),
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

func InitDbMigration() (database2.DbMigration, error) {
	panic(wire.Build(database2.NewDbMigration, CommonWiring, DbWiring))
}

func InitServerTest(ctx context.Context, m *mockBase.Mocks) (srv.Server, error) {
	panic(wire.Build(
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
