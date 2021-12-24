package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan string)
	joe := boring("Joe", quit)

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	fmt.Printf("I'm listening %d times.\n", n)

	// Listen a random numer of times than ask to stop talking
	for ; n > 0; n-- {
		fmt.Printf("%s says\n", <-joe)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string { // a receive only channel
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				cleanup()
				quit <- "See you!"
				return
			}
		}
	}()
	return c
}

func cleanup() {
	fmt.Println("Gonna stop! Cleaning up.")
}
