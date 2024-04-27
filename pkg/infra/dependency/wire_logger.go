//go:build wireinject
// +build wireinject

package dependency

import (
	logger "github.com/Medzoner/medzoner-go/pkg/infra/logger"

	"github.com/google/wire"
)

var (
	LoggerWiring = wire.NewSet(
		logger.NewLoggerAdapter, wire.Bind(new(logger.ILogger), new(*logger.ZapLoggerAdapter)),
	)
)

func initLogger() applogger.CommonLogger {
	panic(wire.Build(wire.NewSet(
		logger.NewLoggerAdapter, wire.Bind(new(logger.ILogger), new(*logger.ZapLoggerAdapter)),
	), InitConfig))
}
