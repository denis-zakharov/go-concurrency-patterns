package main

import (
	"fmt"
	"io"
	"log"
)

var partner = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser) {
	fmt.Fprint(c, "Waiting for a partner...")
	select {
	// Case 1: one user, the partner channel is empty, then
	// put own connection into the channel.
	// Case 2: a new connection comes up, there is a value in the channel.
	// Cannot put a connection into the channel therefore get the existing
	// connection from the channel and host a chat with two connections
	// (own and partner's).
	case partner <- c:
		// subseed to another goroutine
	case p := <-partner:
		// host a chat
		chat(p, c)
	}
}

func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "Found one! Say hi.")
	fmt.Fprintln(b, "Found one! Say hi.")
	// handle disconnects
	errc := make(chan error)
	go cp(a, b, errc)
	go cp(b, a, errc)
	if err := <-errc; err != nil {
		log.Println(err)
	}
	a.Close()
	b.Close()
}

func cp(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(w, r)
	errc <- err
}
