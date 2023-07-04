package event

// IEventHandler is an interface that contains method Handle
type IEventHandler interface {
	Handle(event Event)
}
