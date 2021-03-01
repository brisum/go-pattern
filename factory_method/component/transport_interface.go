package component

type TransportInterface interface {
	CreateMessage() (MessageInterface, error)
}
