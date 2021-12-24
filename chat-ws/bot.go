package main

import (
	"io"
	"time"
)

type bot struct {
	io.ReadCloser
	out io.Writer
}

func NewBot() io.ReadWriteCloser {
	r, out := io.Pipe()
	return bot{r, out}
}

func (bot bot) Write(b []byte) (int, error) {
	go bot.speak()
	return len(b), nil
}

func (bot bot) speak() {
	time.Sleep(1 * time.Second)
	msg := chain.Generate(10)
	bot.out.Write([]byte(msg))
}
