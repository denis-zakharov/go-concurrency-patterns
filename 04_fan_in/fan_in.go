package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fanIn(inputs ...<-chan string) <-chan string {
	c := make(chan string)
	for i := range inputs {
		input := inputs[i] // New instance of 'input' for each loop.
		go func() {
			for {
				c <- <-input // get a value from input and send it to the channel c
				// `c <- input1` sends a channel to a channel
			}
		}()
	}
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	fmt.Println("I'm listening.")
	for i := 0; i < 20; i++ {
		fmt.Printf("%s says\n", <-c)
	}
	fmt.Println("You're boring, I'm leaving.")
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
