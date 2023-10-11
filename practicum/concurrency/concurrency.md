# Concurrency in Go

## Goroutines

"A goroutine is a _lightweight thread of execution_" managed by the go runtime.

To invoke a function in a goroutine, use the `go ...` syntax.

```go
func printAllMultiplesTo100(factor int) {
	for i := 0; i <= 100; i += factor {
		time.Sleep(5 * time.Millisecond)
		fmt.Printf("%d: %d\n", factor, i)
	}
}

// ...

go printAllMultiplesTo100(5)
go printAllMultiplesTo100(10)
printAllMultiplesTo100(3)
```

## Channels

