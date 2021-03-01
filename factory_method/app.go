package main

import component "./component"

type App struct {
}

func (app App) run(transport component.TransportInterface) {
	var message, _ = transport.CreateMessage()
	var _ = message.Send("bear", "Spring is coming", "Wake up!")
}
