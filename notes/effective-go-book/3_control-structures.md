# Effective Go: Control Structures

## Table of Contents

- [If](#if)
- [Redeclaration and reassignment](#redeclaration-and-assignment)
- [For](#for)
- [Switch](#switch)
- [Type switch](#type-switch)

## If

Go `if` statements do not need parentheses, and can include initializations of local variables. Otherwise, they are the same as in C.

```go
if userStatus := user.Status(); userStatus == "online" {
    // do something if true
} else {
    // do something if false
    // can use the local userStatus variable in the else block
}
```

## Redeclaration and Reassignment

In Go, `:=` is the short form of _initialization_ (declaring and assigning a value). You can reassign a variable with `:=` provided that:

- the declaration is in the same scope as the existing declaration
- the value in the assignable (of the static type)
- there is at least one other variable that is created by the `:=` syntax

```go
name := "Ben"

// user.Details returns two values, status and name
// since name is already declared, := is used here to reassign it
status, name := user.Details()
```

## For

In Go, there is not `do/while` or `while` loop. The `for` loops is used in several variations to cover any range of iterations.

Here are common loop structures and their implementation in Go:

- Traditional `for` loop

  ```go
  for i := 0; i < 10; i++ {
      fmt.Println(i)
  }
  ```

- Traditional `while` loop

  ```go
  i := 0
  for i < 10 {
    fmt.Println(i)
    i++
  }
  ```

A `range` clause is useful for managing arrays, slices, strings, maps etc. It is similar to `for/in` and `for/of` (JS) or `foreach` (C#).

```go
hobbits := []string{"Bilbo", "Frodo", "Sam"}

for key, value := range hobbits {
    fmt.Println(value + " at key " + key)
}
```

If you only need the key, omit the second variable. If you only need the value, use a _blank identifier_, denoted by an `_` underscore character.

The `range` will help with iterating over strings. It breaks each character into a Unicode code point, or _rune_.

```go
name := "Frodo"

for _, char := range name {
    fmt.Println(string(char)) // converts unicode code point (rune) to string
}
```

## Switch

In Go, `switch` structures will test a case until a match is found. If no expression is given, the `switch` checks for `true` matches. There are also no _fallthroughs_ (normally prevented by `break` statements), so if you need a fallthrough, commas can be used to combine cases.

```go
month := 7

switch month {
    case 1, 2, 11, 12:
        fmt.Println("Winter")
     case 3, 4, 5:
        fmt.Println("Spring")
    case 6, 7, 8:
        fmt.Println("Summer")
    case 9, 10:
        fmt.Println("Fall")
}
```

## Type Switch

Switches can be used to find the _dynamic type_ of a variable interface. Below is the example from the book:

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```
