package component

type MessageInterface interface {
	Send(destination string, subject string, text string) error
}

type TransportInterface interface {
	CreateMessage() (MessageInterface, error)
}
