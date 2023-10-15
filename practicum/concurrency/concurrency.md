# Concurrency in Go

## Goroutines

"A goroutine is a _lightweight thread of execution_" managed by the go runtime.

To invoke a function in a goroutine, use the `go ...` syntax.

```go
func sum(arr []int) {
	result := 0
	for _, val := range arr {
		result += val
	}
	fmt.Println(result)
}

// ...

go sum([]int {1, 2, 3, 4, 5})
go sum([]int {2, 3, 4, 5, 6})
```

## Channels

_Channels_ are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

Create a channel with the `make(chan type)` function and providing the channel data type.

```go
c := make(chan int)
```

Channels are used to send and receive values. Use the `<-` syntax to denote sent values. We'll change up the `sum` function to take a channel as a parameter and send the sum value through.

```go
func sum(arr []int, c chan int) {
	result := 0
	for _, val := range arr {
		result += val
	}
	c <- result // sends result into the channel
}
```

Then the same syntax is used to receive a value from a channel.

```go
go sum([]int{1, 2, 3, 4, 5}, c)
go sum([]int{2, 3, 4, 5, 6}, c)

result1, result2 := <-c, <-c
```

By default, sending and receiving values in/from a channel is a blocking process so that no explicit goroutine synchronization is required.

### Channel Buffering

By default, channels are also _unbuffered_, meaning they will only accept sent values if there are corresponding receivers. _Buffered_ channels are limited to a predefined number of sent/received values, and can even receive values not sent from concurrency.

```go
c := make(chan int, 2)

c <- 1
c <- 2

value1 := <- c
value2 := <- c
```

### Channel Synchronization

Channel receivers are blocking by default. This can allow you to synchronize your goroutines by blocking a program until a goroutine is finished.

```go
func sum(arr []int, c chan bool) {
	result := 0
	for _, val := range arr {
		result += val
	}
	fmt.Println(result)
	c <- true
}

// ...

c := make(chan bool)

go sum([]int {1, 2, 3, 4, 5}, c)

<- c // blocks program from exiting until sum goroutine is finished
```

### Channel Directions

When passing channels as function parameters, you can specify the channel direction (sending or receiving) by adding the `<-` syntax before or after the type declaration in the function definition.

```go
func in(c chan<- string, msg string) {
	c <- msg
}

func out(c <-chan string) {
	msg := <- c
	fmt.Println(msg)
}

// ..

messenger := make(chan string, 1)
in(messenger, "Hello World!")
out(messenger)
```

### Using `select` with Channels

The `select` statement lets a goroutine wait on multiple communication operations.

The default case can be used to run code when a channel is not ready to receive yet.

```go
select {
	case value1 := <-channel1
		fmt.Println(value1)
	case value2 := <-channel2
		fmt.Println(value2)
	default
		fmt.Println("waiting on channels")
}
```

### Range and Closing Channels

You can use `range` to iterate over the unread values in a buffered channel. You can also `close` a channel to signify to the receiver that there are no more values to be read from the channel.

```go
func waitForMe(c chan int, length int) {
	for i := 1; i <= length; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(c)
}

// ...

queue := make(chan int, 3)
go waitForMe(queue, 3)

for value := range queue {
	fmt.Println(value)
}
```

Here, the range will iterate over the values in the buffered channel until it `close(chan)` is called on the sending side.

## Timeouts, Timers, and Tickers

### Timeouts

_Timeouts_ are easy and useful in go programs to avoid memory leaks. Since `select` structures are useful in handling multiple concurrent operations, they can be further used to manage operations with defined timeouts.

```go
func waitForMe(c chan string, seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
	c <- "done!"
}

waiter := make(chan string, 1)

go waitForMe(waiter, 3)

select {
case msg := <-waiter:
	fmt.Println(msg)
case <-time.After(2 * time.Second):
	fmt.Println("timeout")
}
```

Here, the `select` runs the case that is ready first, either the waiter channel or the 2 second timeout.

### Timers

_Timers_ are useful when you want some code to execute once in the future after a designated amount of time.

Create a timer with the `time.NewTimer(d duration)` syntax. This will provide a channel that will be notified after the duration has completed.

```go
timer := time.NewTimer(time.Second)

<- timer.C // channel that is notified after one second (is blocking)
```

Timers can be stopped with the `timer.Stop()` method. This differentiates timers from `time.Sleep`.

### Tickers

_Tickers_ are similar to timers, but repeat after a given interval duration.

Create a ticker with the `time.NewTicker(d duration)` syntax. This will provide a channel that will be notified each time the ticker hits the interval.

```go
ticker := time.NewTicker(time.Second)
timer := time.NewTimer(5 * time.Second)

go func() {
	for tick := range ticker.C {
		fmt.Println(tick)
	}
}()

<-timer.C
ticker.Stop()
```

Tickers can be stopped the same way _timers_ are stopped with the `ticker.Stop()` method.
