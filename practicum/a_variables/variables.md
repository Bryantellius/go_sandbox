# Variables and Values in Go

## Values

Go has basic data types that cover a host of scenarios.

| Type       | Description                                                                      | Example                               |
| ---------- | -------------------------------------------------------------------------------- | ------------------------------------- |
| bool       | a logical boolean value                                                          | `true` or `false`                     |
| string     | a collection of text characters                                                  | `"Hello World!"`                      |
| int        | integer value                                                                    | `1`                                   |
| uint       | unsigned integer value                                                           | `12`                                  |
| uintptr    | unsigned integer large enough to store the uninterpreted bits of a pointer value | `123`                                 |
| int8       | integer up to 8 bits                                                             | `127` (max value)                     |
| uint8      | unsigned integer up to 8 bits                                                    | `255` (max value)                     |
| int16      | integer up to 8 bits                                                             | `32767` (max value)                   |
| uint16     | unsigned integer up to 8 bits                                                    | `65535` (max value)                   |
| int32      | integer up to 32 bits                                                            | `2147483647` (max value)              |
| uint32     | unsigned integer up to 32 bits                                                   | `4294967295` (max value)              |
| int64      | integer up to 64 bits                                                            | `9223372036854775807` (max value)     |
| uint64     | unsigned integer up to 64 bits                                                   | a large number                        |
| byte       | same as uint8                                                                    | `255` (max value)                     |
| rune       | same size as uint32, represents a unicode code point                             | `'B'`                                 |
| float32    | floating point number up to 32 bits                                              | `3.4028234663852886e+38` (max value)  |
| float64    | floating point number up to 64 bits                                              | `1.7976931348623157e+308` (max value) |
| complex64  | complex number up to 64 bits                                                     | `3 + 4i`                              |
| complex128 | complex number up to 128 bits                                                    | `11 + 6i`                             |

## Variables

Variables are explicitly declared with the `var` keyword. They can be declared inside or outside of functions.

You can declare one or more variables with the `var` keyword, with the variable names and values separated by commas, _as long as the variables hold the same types_. Otherwise, you can declare multiple variables together in a _factored declaration_.

Since Go is a statically typed language, variables are declared with the data type after the variable name. Go also has inferred typing, so if no data type is present after a variable name in a declaration, Go will _infer_ the type for the variable based on the assigned value at declaration.

You can assign a value to a variable at the time of declaration with the `=` symbol followed by a value literal. Variable values can also be _reassigned_ new values after declaration with the same symbol.

Lastly, you can use the `:=` symbol as shorthand to declare and assign a variable value _within a function_.

```go
// declaring a variable of type string outside a function
var name string

// declaring a variable of type string outside a function and assigning a value
var birthMonthStr string = "January"

// declaring a variable with inferred typing inside a function and assigning a value
var birthMonthNum = 1 // inferred to be a integer

// declare multiple variables at once
var a, b, c int = 1, 2, 3

// factored declaration
var (
	age          int = 31
	isActiveUser     = true
	username         = "rs3"
)

func PrintHello() {
    // declaring a variable of type string inside a function
    var message string

    // declaring a variable of type int (integer) inside a function and assigning a value
    var favNumber int = 5

    // declaring a variable with inferred typing inside a function and assigning a value
    var ten = 10 // inferred to be a integer

    // declaring and assigning a variable with inferred typing inside a function using shorthand
    isOnline := true
}
```

### Constant

Some variables can contain constant values. Use the `constant` keyword to declare a variable and assign it a constant value.

> Constant values cannot be reassigned within a program, can be explicitly typed or inferred, and must be assigned values at declaration. Constant values can be characters, strings, booleans or numeric data types.

```go
// constant variable with explicit type
const daysInYear float64 = 365.25

// constant variable with inferred type
const january = "January"
```

### Naming Conventions for Variables

Go naming conventions follow a few simple rules:

- _exported names_ are written in _Pascal Case_ (`ThisIsPascalCase`)
- all other names are written in _camel case_ (`thisIsCamelCase`)
- character limitations are similar to other languages
  - can include numbers, characters, or `_`
  - cannot start a name with a number
- all variables within a function are camel case (they are within the scope of a function and cannot be exported)

### Default Values

Variables that are not declared with a value are assigned _zero_ values based on their explicit data type.

- strings `""`
- numbers `0`
- boolean `false`

### Type Conversion

Values can be converted to other data types, but must be done _explicitly_ using built-in functions. Each basic type has a conversion function.

```go
// convert an int to a float64
var radius int = 11
const PI float64 = 3.14

area := PI * float64(radius*radius)
```

Go does not have _implicit type conversion_, so the following would not work.

```go
var age int = 20
fmt.Println("He is " + age + " years old")
```
