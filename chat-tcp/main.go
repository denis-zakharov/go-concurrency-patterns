package main

import (
	"fmt"
	"log"
	"net"
)

const listenAddr = "localhost:40000"

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept() // net.Conn implements io.Writer
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(c, "Hello!")
		go match(c)
	}
}
