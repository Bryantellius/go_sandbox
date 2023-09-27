# Effective Go: Functions

## Table of Contents

- [Multiple return values](#multiple-return-values)
- [Named result parameters](#named-result-parameters)
- [Defer](#defer)

## Multiple Return Values

Go functions can have multiple return values, seperated by _commas_.

```go
func checkAnswer(answer string) (string, bool) {
    if (answer == "blue") {
        return "blue", true
    } else {
        return "blue", false
    }
}
```

## Named Result Parameters

_Named results_ are named return values. Although not necessary, naming result parameters can be good a documentation practice.

```go
func Add(x int, y int) (sum int) {
    return x + y
}
```

## Defer

The `defer` keyword can be used to _defer_ calling a function until immediately _before_ the calling function closes.

Deferred functions are called in a LIFO order. Take the following example:

```go
for i := 0; i <= 10; i++ {
    defer fmt.Println(i)
}
```

In a mind-bending example, take note of when deferred function arguments are evaluated:

```go
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```

The above example would print:

```txt
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```
