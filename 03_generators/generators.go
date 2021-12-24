package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	fmt.Println("I'm listening.")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s says\n", <-joe)
		fmt.Printf("%s says\n", <-ann)
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
