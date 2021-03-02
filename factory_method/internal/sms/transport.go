package sms

import component "../component"

type Transport struct {
}

func (transport Transport) CreateMessage() (component.MessageInterface, error) {
	mailMessage := Message{}

	return mailMessage, nil
}
