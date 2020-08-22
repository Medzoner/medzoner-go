package messagebus

import (
	"context"
	"github.com/Medzoner/medzoner-go/pkg/application/utils/messager"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"time"
)

type CommandBus struct {
	Bus *cqrs.CommandBus
	Handlers []cqrs.CommandHandler
	Router *message.Router
}

func (c *CommandBus) NewBus() messager.MessageBus {
	logger := watermill.NewStdLogger(false, false)
	cqrsMarshaler := cqrs.JSONMarshaler{}

	commandsPublisher := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	routerCommandBus, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	routerCommandBus.AddMiddleware(middleware.Recoverer)

	cqrsFacade, err := cqrs.NewFacade(cqrs.FacadeConfig{
		GenerateCommandsTopic: func(commandName string) string {
			return commandName
		},
		CommandHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
			return c.Handlers
		},
		CommandsPublisher: commandsPublisher,
		CommandsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			return commandsPublisher, nil
		},

		Router:                routerCommandBus,
		CommandEventMarshaler: cqrsMarshaler,
		Logger:                logger,
	})
	if err != nil {
		panic(err)
	}

	c.Bus = cqrsFacade.CommandBus()
	c.Router = routerCommandBus
	return c
}

func (c *CommandBus) Run()  {
	if err := c.Router.Run(context.Background()); err != nil {
		panic(err)
	}
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
}

func (c *CommandBus) Handle(message messager.Message)  {
	c.publishCommands(message)
}

func (c *CommandBus) publishCommands(msg messager.Message) {
	cmd := &msg
	if err := c.Bus.Send(context.Background(), cmd); err != nil {
		panic(err)
	}
}
