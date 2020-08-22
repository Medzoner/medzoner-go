package messager

type MessageBus interface {
	Handle(message Message)
	NewBus() MessageBus
}

type Message interface {
	GetName() string
}
