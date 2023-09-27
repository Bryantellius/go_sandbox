# Effective Go: Methods

## Table of Contents

- [Pointers vs. Values](#pointers-vs-values)

## Pre-requisites

Methods can be defined for any _named_ type besides pointers or interfaces. You can define a method for a type as in the following example:

```go
type Hobbit struct {
    name string
    age  int
}

func (h Hobbit) Details() {
    fmt.Printf("%v is %d years old", h.name, h.age)
}

func main() {
    var frodo Hobbit = Hobbit{"Frodo", 51}
    frodo.Details()
}
```

## Pointers vs Values

"The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers."

Pointers are _addresses_ to a value in memory, whereas values are the _values_ themselves. If we want a method to mutate the caller's data, we must create a pointer method that will act on a pointer to a value.

Here is _Effective Go_'s example:

```go
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // Again as above.
    *p = slice
    return len(data), nil
}
```
