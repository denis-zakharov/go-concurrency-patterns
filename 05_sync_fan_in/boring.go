package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	fmt.Println("I'm listening.")
	for i := 0; i < 20; i++ {
		msg1 := <-c
		fmt.Printf("%s says\n", msg1.str)

		msg2 := <-c
		fmt.Printf("%s says\n", msg2.str)

		msg1.wait <- true // confirm to producer 1
		msg2.wait <- true // confirm to producer 2
	}
	fmt.Println("You're boring, I'm leaving.")
}

func boring(msg string) <-chan Message { // a receive only channel
	c := make(chan Message)
	waitForIt := make(chan bool) // Shared between all messages of this producer
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt // wait until consumed and confirmed
		}
	}()
	return c
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
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
