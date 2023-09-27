# Effective Go: Interfaces and Other Types

## Table of Contents

- [Interfaces](#interfaces)
- [Conversions](#conversions)
- [Interface conversions and type assertions](#interface-conversions-and-type-assertions)
- [Generality](#generality)
- [Interfaces and methods](#interfaces-and-methods)
- [Embedding](#embedding)

## Interfaces

Interfaces define the structure of an object. In Go, it's common to have interfaces with only one method, making the interface name a verb of the method name (ie `Writer` that has a single `Write`) method. A type can implement multiple interfaces.

Here's an example of a type with a method:

```go
type Hobbits []string

func (h Hobbits) Introduce() {
    res := ""

    for i, hobbit := range h {
        if i == len(h) - 1 {
            res += " and " + hobbit
        } else {
            res += " " + hobbit
        }
    }

    fmt.Printf("Here are%v", res)
}
```

`Hobbits` can implement the `sort` methods if it adheres to the `sort.Interface`. We'll need to add `Len`, `Less`, and `Swap` methods.

```go
func (h Hobbits) Len() int {
    return len(h)
}
func (h Hobbits) Less(i, j int) bool {
    return h[i] < h[j]
}
func (h Hobbits) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
```

## Conversions

Data types can be converted from one to another. For example, integers can be converted to float64 and vice versa with `int(4.02)` or `float64(4)`. This also goes for `[]string`, `[]int`, etc.

```go
party := Hobbits{"Bilbo", "Frodo", "Thorin", "Samwise" }
[]string(party) // converts the Hobbits type ([]string) to a regular []string

someNumber := 4.5
int(someNumber) // returns 4
```

## Interface Conversions and Type Assertions

Use a _type switch_ in go to deterministically convert data types.

```go
type Stringer interface {
    String() string
}

var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```

A _type assertion_ can be used to extract a string value from a type by name.

```go
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
```

The `ok` value is used to ensure that the value contains a string.

## Generality

"If a type exists only to implement an interface and will never have exported methods beyond that interface, there is no need to export the type itself. Exporting just the interface makes it clear the value has no interesting behavior beyond what is described in the interface. It also avoids the need to repeat the documentation on every instance of a common method." - [Essential Go, Generality](https://go.dev/doc/effective_go#formatting)

## Interfaces and Methods

Most types can have methods defined for them, which means that most types can satisfy interfaces.

For example, the `http.Handler` interface looks as such:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

We can make regular functions handle http requests by adding methods to types to have them satisfy the Handler interface.

```go
import "net/http"

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.  If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, req).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
    f(w, req)
}

// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, os.Args)
}

// ArgServer is converted to a HandlerFunc interface and used
http.Handle("/args", http.HandlerFunc(ArgServer))
```

Because of this, you can view interfaces as a set of methods, which can be defined for almost any type.

## Embedding

In Javascript and other languages, classes can have _parent_ or _super_ classes that they inherit from. In Go, this is a similar yet contrast concept called _embedding_.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
```

Embedding rules:

- Nested fields are _dominated_ by top level fields of the same name (similar to method/field _overriding_ in other languages)
- Duplicate names with cause errors, unless the names are never used
