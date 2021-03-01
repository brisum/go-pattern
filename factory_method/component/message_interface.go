package component

type MessageInterface interface {
	Send(destination string, subject string, text string) error
}
