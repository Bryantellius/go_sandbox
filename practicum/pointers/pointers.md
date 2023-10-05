# Pointers in Go

_Pointers_ are memory addresses to values with a program. They allow you to access and pass _reference_ values.

A pointer is signified by the `*` operator before the variable.

```go
func ToUpper(name *string) { // specifies that the parameter will be a pointer of type string
    *name = strings.ToUpper(*name) // dereferences the pointer by reassigning to a new value
}
```

> When a pointer value (address) is changed from it's previous value, it is referred to as _dereferencing_.

The `&` operator before a variable generates a pointer to the operand.

```go
var name string = "Ben"
fmt.Println(name) // "Ben"
ToUpper(&name)  // passes a pointer to the name value
fmt.Println(name) // "BEN"
```

> Data structures, such as arrays, slices, maps and structs, are already passed by reference.
