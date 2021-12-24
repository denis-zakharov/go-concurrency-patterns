package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

const listenAddr = "localhost:40000"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, listenAddr)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/", websocket.Handler(socketHandler))
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
