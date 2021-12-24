### Go Concurrency Patterns by Rob Pike
- [Video](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Slides](https://talks.golang.org/2013/advconc.slide)
- [Code](https://talks.golang.org/2012/concurrency/support)

### Advanced Concurrency Patterns by Sameer Ajmani
- [Video](https://www.youtube.com/watch?v=QDDwwePbDtw)
- [Slides](https://talks.golang.org/2013/advconc.slide)

### Share Memory by Communicating
- [Codewalk](https://go.dev/doc/codewalk/sharemem/)

### Chat-Roulette
A toy example of concurrency usage.
- [Video](https://vimeo.com/53221560)
- [Slides](https://talks.golang.org/2012/chat.slide)

## Go Concurrency Patterns TOC
### 01_goroutines
A basic goroutine for asynchronous running.

### 02_channels
A basic channel (non-buffered) to communicate and synchronize.

```go
var c = make(chan string)

func consumer(receiveOnlyChannel <-chan string) {
    // consume
    fmt.Println(<-receiveOnlyChannel)
}
```

### 03_generator
A function that returns a channel. A return type `<-chan string` means a receive only channel of strings.

We can generate several channels and then iterate over them in order.

### 04_fan_in
Multiplex several channels into one to send the messages as soon as they are ready (that is
not necessarily in order).

```
ch1 ->
      \
      multiplex -->
      / 
ch2 ->
```

### 05_sync_fan_in
We can restore sequencing in the fan-in pattern by using private channels shared between
all messages *of each producer*.

Receive messages from all producers, then enable them again by sending back on a private channel.

### 06_select
The select statement provides another way to handle multiple channels.

It's like a switch, but each case is a communication:
- All channels are evaluated.
- Selection blocks until one communication can proceed, which then does.
- If multiple can proceed, select chooses pseudo-randomly.
- A default clause, *if present*, executes immediately if no channel is ready.

### 07_select_timeout
The `time.After` function returns a channel that blocks for the specified duration.
After the interval, the channel delivers the current time, once.

### 08_quit_channel
Tell the producer to stop 'talking'.

### 09_daisy_chain

Create a chain of deferred computations using channels.

A part of `daisy_chain.go` example for `n=3`:

```
[leftmost] <-- 1 + [r1] <-- 1 + [r2] <-- 1 + [r3]
                                               ^
                                               |___init value 0

leftmost <- 1 + (r1 <- 1 + (r2 <- 1 + r3))
leftmost <- 1 + (r1 <- 1 + (r2 <- 1 + 0))
leftmost <- 1 + (r1 <- 1 + 1)
leftmost <- 1 + 2
<- leftmost // result = 3
```

### 10_system_interactions

Let us construct the naive fake Google search:
- send requests to search subsystems
  - asyncronous
  - replicated
  - with timeout
- mix the responses in some way
- respond to a user

### summary
Goroutines and channels are big ideas. They're tools for program construction.

But sometimes all you need is a reference counter.

Go has "sync" and "sync/atomic" packages that provide mutexes, condition variables, etc.
They provide tools for smaller problems.

Often, these things will work together to solve a bigger problem.