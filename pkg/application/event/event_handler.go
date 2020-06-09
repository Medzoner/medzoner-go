package event

type EventHandler interface {
	Handle(event Event)
}
