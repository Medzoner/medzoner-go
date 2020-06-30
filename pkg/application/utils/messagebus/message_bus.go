package messagebus

type MessageBus interface {
	Handle(message Message)
}

type Message interface {
	GetName() string
}
