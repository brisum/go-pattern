package sms

import (
	"fmt"
)

type Message struct {
}

func (message Message) Send(destination string, subject string, text string) error {
	fmt.Println("Send sms message...")
	fmt.Printf("destination %v\n", destination)
	fmt.Printf("subject %v\n", subject)
	fmt.Printf("text %v\n", text)

	return nil
}
