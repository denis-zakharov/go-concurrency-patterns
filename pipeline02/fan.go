package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)
	for i := range cs {
		gen := cs[i]
		var g = i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range gen {
				fmt.Printf("[%d] consumed by GO[%d]\n", n, g)
				out <- n
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Fan-out, Fan-in or Scatter-Gather
func main() {
	in := gen(2, 3, 5)

	// Fan-out: distribute workload accross two sq goroutines.
	// This is similar to a consumer group in Kafka.
	c1 := sq(in)
	c2 := sq(in)

	// Fan-in: consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}
