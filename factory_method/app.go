package main

import (
	component "./internal/component"
	mail "./internal/mail"
	sms "./internal/sms"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var factory = component.NewFactory()
	factory.Register(component.TRANSPORT_EMAIL_TYPE, reflect.TypeOf(mail.Transport{}))
	factory.Register(component.TRANSPORT_SMS_TYPE, reflect.TypeOf(sms.Transport{}))

	var transportType string
	var transport component.TransportInterface
	var message component.MessageInterface

	if rand.Intn(2) == 1 {
		transportType = component.TRANSPORT_SMS_TYPE
	} else {
		transportType = component.TRANSPORT_EMAIL_TYPE
	}

	transport, err := factory.CreateMessageTransport(transportType)
	fmt.Printf("Transport: %T\n", transport)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	message, err = transport.CreateMessage()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = message.Send("bear", "Spring is coming...", "Wake up!")
}
