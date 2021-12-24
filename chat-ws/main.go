package main

import (
	"log"
	"net"
	"net/http"

	"golang.org/x/net/websocket"
)

const tcpListenAddr = "localhost:4000"
const wsListenAddr = "localhost:40000"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, wsListenAddr)
}

func netListen() {
	l, err := net.Listen("tcp", tcpListenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go match(c)
	}
}

func main() {
	go netListen()
	http.HandleFunc("/", rootHandler)
	http.Handle("/socket", websocket.Handler(socketHandler))
	if err := http.ListenAndServe(wsListenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
