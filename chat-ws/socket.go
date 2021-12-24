package main

import (
	"io"

	"golang.org/x/net/websocket"
)

type socket struct {
	io.ReadWriter
	done chan bool
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

func socketHandler(ws *websocket.Conn) {
	// We can't just use a `websocket.Conn` for a chat,
	// because the `ws` is held open by its handler function.
	// Here we use a channel to keep the handler running
	// until the socket's Close method is called.
	s := socket{ws, make(chan bool)}
	go match(s)
	<-s.done // block until the chat end (s.Close() in a chat)
}