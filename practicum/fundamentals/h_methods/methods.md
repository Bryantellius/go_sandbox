# Methods in Go

Go does not have classes, but it does still support methods on types (structs or custom types). In Go, _methods_ are just functions with _receiver types_.

Receiver types can either by _pointer_ or _values_ types. They are declared between the `func` keyword and the function name.

Methods can only be declared in packages where their receiver types are declared.

```go
type person struct {
    name string
}

func (p person) sayHello(name string) {
    fmt.Printf("Hello %s! My name is %s.", name, p.name)
}
```

Above, `sayHello` is a method for the `person` struct type. It receives a _value_ receiver type of `p person`, designating it to be a method for that type.

```go
type bankAccount struct {
    accountHolder string
    balance       float64
}

func (ba *bankAccount) reportTransaction(amount float64) float64 {
    ba.balance += amount
    return ba.balance
}
```

Here we have a _pointer receiver type_ defined on the `reportTransaction` method. This allows us to mutate the receiver value within the function through the pointer.
