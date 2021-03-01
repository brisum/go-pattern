package main

import (
	component "./component"
	mail "./mail"
	sms "./sms"
	"math/rand"
	"time"
)

func main() {
	var transport component.TransportInterface
	var app App = App{}

	rand.Seed(time.Now().UnixNano())

	if rand.Intn(2) == 1 {
		transport = sms.Transport{}
	} else {
		transport = mail.Transport{}
	}

	app.run(transport)
}
