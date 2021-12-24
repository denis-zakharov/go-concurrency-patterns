package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("Joe")
	fmt.Println("I'm listening.")
	timeout := time.After(10 * time.Second) // global conversation timeout
	for {
		select {
		case s := <-joe:
			fmt.Printf("%s says\n", s)
		case <-time.After(890 * time.Millisecond): // timeout for each message
			fmt.Println("Joe's been too slow.")
			return
		case <-timeout:
			fmt.Println("Joe's talked too much.")
			return
		}
	}
}

func boring(msg string) <-chan string { // a receive only channel
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
