package command_bus

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/utils/messagebus"
)
import (
	"github.com/mustafaturan/bus"
	"github.com/mustafaturan/monoton"
	"github.com/mustafaturan/monoton/sequencer"
)

func (c CommandBus) NewBus() *bus.Bus {
	// configure id generator (it doesn't have to be monoton)
	node        := uint64(1)
	initialTime := uint64(1577865600000) // set 2020-01-01 PST as initial time
	m, err := monoton.New(sequencer.NewMillisecond(), node, initialTime)
	if err != nil {
		panic(err)
	}

	// init an id generator
	var idGenerator bus.Next = (*m).Next

	// create a new bus instance
	c.Bus, err = bus.NewBus(idGenerator)
	if err != nil {
		panic(err)
	}

	// maybe register topics in here
	c.Bus.RegisterTopics("order.received", "order.fulfilled")

	return c.Bus
}

type CommandBus struct {
	Bus *bus.Bus
}

func (c *CommandBus) Handle(message messagebus.Message)  {
	ctx := context.Background()
	ctx = context.WithValue(ctx, bus.CtxKeyTxID, "some-transaction-id-if-exists")

	event, err := c.Bus.Emit(ctx, message.GetName(), message)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(event)
}
