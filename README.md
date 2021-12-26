# Go Concurrency Patterns by Rob Pike
- [Video](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Slides](https://talks.golang.org/2013/advconc.slide)
- [Code](https://talks.golang.org/2012/concurrency/support)

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

## summary
Goroutines and channels are big ideas. They're tools for program construction.

But sometimes all you need is a reference counter.

Go has "sync" and "sync/atomic" packages that provide mutexes, condition variables, etc.
They provide tools for smaller problems.

Often, these things will work together to solve a bigger problem.

# Chat-Roulette by Andrew Gerrand
A toy example of concurrency usage.
- [Video](https://vimeo.com/53221560)
- [Slides](https://talks.golang.org/2012/chat.slide)

# Advanced Concurrency Patterns by Sameer Ajmani

- [Video](https://www.youtube.com/watch?v=QDDwwePbDtw)
- [Slides](https://talks.golang.org/2013/advconc.slide)

An example how to merge several subsriptions into one `feed` with cancellation.

# Pipelines by Sameer Ajmani

- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)

`pipeline_01` is a simple example how to chain computations of square numbers.

`pipeline_02` is a combination of the fan-in and fan-out patterns into a pipeline.

`pipeline_03` is an example of early stopping - producers stop when consumer wants to stop earlier.

`pipeline_04` is an example of a bounded parallelism: spawn a goroutine to read a file and
calculate md5 sum for it in a directory tree with un upper bound on number of files/goroutines.

The guidelines for pipeline construction:

- stages close their outbound channels when all the send operations are done.
- stages keep receiving values from inbound channels until those channels are closed or the senders are unblocked.

# Context by Sameer Ajmani

- [Go Concurrency Patterns: Context](https://go.dev/blog/context)

```go
// A Context carries a deadline, cancellation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```

The `Done` method returns a channel that acts as a cancellation signal to functions running on behalf
of the `Context`: when the channel is closed, the functions should abandon their work and return.

The `Err` method returns an error indicating why the `Context` was canceled.

A `Context` does *not* have a Cancel method for the same reason the Done channel is receive-only:
the function receiving a cancellation signal is usually not the one that sends the signal. In particular,
when a parent operation starts goroutines for sub-operations, those sub-operations should not be able to cancel the parent. Instead, the `WithCancel` function provides a way to cancel a new `Context` value.

A Context is safe for simultaneous use by multiple goroutines. Code can pass a single `Context`
to any number of goroutines and cancel that `Context` to signal all of them.

The `Deadline` method allows functions to determine whether they should start work at all; if too little
time is left, it may not be worthwhile. Code may also use a deadline to set timeouts for I/O operations.

`Value` allows a `Context` to carry request-scoped data. That data must be safe for simultaneous
use by multiple goroutines.

**Context tree**

```
func Background() Context // never canceled
|
|\____func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
|
|\____func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
|
 \____func WithValue(parent Context, key interface{}, val interface{}) Context
```

# Share Memory by Communicating
- [Codewalk](https://go.dev/doc/codewalk/sharemem/)

