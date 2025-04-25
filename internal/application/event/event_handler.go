package event

import "context"

// IEventHandler is an interface that contains method Handle
type IEventHandler interface {
	Publish(ctx context.Context, event Event) error
}
