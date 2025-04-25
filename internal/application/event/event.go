package event

// Event is an interface that contains model and event name
type Event interface {
	GetModel() interface{}
}
