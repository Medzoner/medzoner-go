package event

//IEventHandler IEventHandler
type IEventHandler interface {
	Handle(event Event)
}
