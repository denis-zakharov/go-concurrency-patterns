package main

import (
	"fmt"
	"sync"
	"time"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()

	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()

	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)
	for i := range cs {
		gen := cs[i]
		var g = i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range gen {
				select {
				case out <- n:
					fmt.Printf("GOROUTINE[%d]: consumed %d\n", g, n)
				case <-done:
					fmt.Printf("GOROUTINE[%d]: done\n", g)
					return
				}
			}
			fmt.Printf("GOROUTINE[%d]: done because inbound channel is closed\n", g)
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func take(done chan struct{}, out <-chan int, n int) {
	// A receive operation on a closed channel can always proceed immediately,
	// yielding the element type’s zero value.
	defer close(done)

	for v := range out {
		fmt.Println(v)
		n--
		if n == 0 {
			return
		}
	}

}

// Fan-out, Fan-in or Scatter-Gather
func main() {
	done := make(chan struct{})
	in := gen(done, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17)

	// Fan-out: distribute workload accross two sq goroutines.
	// This is similar to a consumer group in Kafka.
	c1 := sq(done, in)
	c2 := sq(done, in)
	c3 := sq(done, in)

	// Fan-in: consume the merged output from c1 and c2.
	take(done, merge(done, c1, c2, c3), 6)

	fmt.Println("Sleeping 3 seconds...")
	time.Sleep(3 * time.Second)
}
