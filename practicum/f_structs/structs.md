# Structs in Go

A _struct_ is a typed collection of fields. They can be created with `type Name struct {...}` syntax.

```go
type runner struct {
    name  string
    place int
}
```

It is common practice in idiomatic Go to encapsulate struct creations in constructors.

```go
func createRunner(name string) *runner {
    newRunner := runner{name: name}
    return &newRunner
}
```

When this is written, the constructor returns a _pointer_ to the created struct for efficiency and reference after the constructor scope is closed.

Struct values are access via _dot notation_.

```go
bb := createRunner("Ben")
hw := createRunner("Hayden")

bb.place = 1
hw.place = 13
```

_Anonymous_ struct types can be used as well, which are typically seen with single value structs.

```go
value := struct {
    mode string
    username string
} {
    "Easy",
    "Bryantellius"
}
```
