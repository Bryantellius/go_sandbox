# Effective Go: Initialization

## Table of Contents

- [Constants](#constants)
- [Variables](#variables)
- [The init function](#the-init-function)

## Constants

In Go, constants are created at _compile-time_, and can be numbers, characters (runes), strings or booleans. Constant variable values cannot be reassigned.

Use a `const` keyword, variable _identifier_, and explicit data type:

```go
const hobbit = "Frodo"

hobbit = "Bilbo" // cannot assign to hobbit (constant "Frodo" of type string)
```

## Variables

Variables are different from constants because they can be initialized at _run-time_.

Use a `var` keyword, variable _identifier_, and explicit data type:

```go
var hobbit = "Frodo"

hobbit = "Bilbo" // can assign to hobbit
```

## The init Function

The `init` function can be defined in each source file. It is called after all variables are initialized. This can be useful for checking program state values upon program initialization.

> Keep scope in mind. The `init` function will only be able to check accessible values within scope.

```go
func init() {
    if hobbit == "" {
        hobbit = "Bilbo"
    }
}
```
