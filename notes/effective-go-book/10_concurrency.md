# Effective Go: Concurrency

## Table of Contents

- [Share by communicating](#share-by-communicating)
- [Goroutines](#goroutines)
- [Channels](#channels)
- [Channels of channels](#channels-of-channels)
- [Parallelization](#parallelization)
- [A leaky buffer](#a-leaky-buffer)

## Share by Communicating

"Do not communicate by sharing memory; instead, share memory by communicating."

Go Routines are designed to share access to variables during concurrency executions without a seperate thread actually accessing the memory. This design also prevents race conditions.

Go's concurrency approach originates from [Hoare's Communicating Sequential Processes (CSP)](https://en.wikipedia.org/wiki/Communicating_sequential_processes), but it is similar to [Unix Pipelines](<https://en.wikipedia.org/wiki/Pipeline_(Unix)>).

## Goroutines

_Goroutines_ are functions executing concurrently within the same address space. Since goroutines are "multiplexed onto multiple OS threads", they are non-blocking.

To run a function as a goroutine, prefix the function call with `go`.

## Channels

_Channels_ are data structures for communication between goroutines.

Create a channel with `make(chan int)`, with an optional second parameter that specifies buffer size (defaults 0).

View Effective Go's simple example:

```go
c := make(chan int)  // Allocate a channel.
// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
<-c   // Wait for sort to finish; discard sent value.
```

## Channels of Channels

Go channels are treated as _first-class citizens/values_, meaning they can be passing around and assigned to variables within a program.

## Parallelization

Parallelization is Go can be achieved to executed expensive operations in parallel across multiple CPU cores.

Effective Go's Example:

```go
// var numCPU = runtime.NumCPU() // number of CPU cores on the machine
var numCPU = runtime.GOMAXPROCS(0) // number of CPU cores available

func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)  // Buffering optional but sensible.
    for i := 0; i < numCPU; i++ {
        go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
    }
    // Drain the channel.
    for i := 0; i < numCPU; i++ {
        <-c    // wait for one task to complete
    }
    // All done.
}
```

Go is a concurrent language, not a parallel one. Although you can model a program to achieve one or the other, or both; they are not the same. Read [Concurrency is not Parallelism](https://go.dev/blog/waza-talk) for more.

## A Leaky Buffer

The following example is from Effective Go on concurrency and buffer management:

```go
var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
    for {
        var b *Buffer
        // Grab a buffer if available; allocate if not.
        select {
        case b = <-freeList:
            // Got one; nothing more to do.
        default:
            // None free, so allocate a new one.
            b = new(Buffer)
        }
        load(b)              // Read next message from the net.
        serverChan <- b      // Send to server.
    }
}

func server() {
    for {
        b := <-serverChan    // Wait for work.
        process(b)
        // Reuse buffer if there's room.
        select {
        case freeList <- b:
            // Buffer on free list; nothing more to do.
        default:
            // Free list full, just carry on.
        }
    }
}
```
