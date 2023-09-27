# Effective Go: Data

## Table of Contents

- [Allocation with new](#allocation-with-new)
- [Constructors and composite literals](#constructors-and-composite-literals)
- [Allocation with make](#allocation-with-make)
- [Arrays](#arrays)
- [Slices](#slices)
- [Two-dimensional slices](#two--dimensional-slices)
- [Maps](#maps)
- [Printing](#printing)
- [Append](#append)

## Allocation with New

Go has a built-in `new` function for creating a _zero-memory address pointer_ of a certain data type.

```go
type Hobbit struct {
    name string
    age  int
}

func main() {
    frodo := new(Hobbit)
    fmt.Println(frodo) // prints &{ 0}
}

```

## Constructors and Composite Literals

A _composite literal_ creates a new instance each time it is used.

```go
type Hobbit struct {
    name string
    age  int
}

func main() {
    frodo := Hobbit{name: "Frodo", age: 51}
    fmt.Println(frodo.name) // prints "Frodo"
}
```

Composite literals can be used for arrays, slices, maps, etc.

## Allocation with make

Go has a built-in `make` function for allocating memory for slices, maps and channels. It takes up to three arguments: type, length, and capacity.

```go
hobbits := make([]string, 10, 100) // creates an slice of length 10 with a capacity of 100
```

## Arrays

Arrays are building blocks for slices, and are treated as values. That means that passing an array value copies it's internal values. You can use a pointer to pass a reference of an array value.

```go
hobbits := [3]string{ "Frodo", "Bilbo", "Samwise" }
```

## Slices

Slices are used more often than arrays in Go. Slices hold references to underlying arrays. Since arrays are statically sized, appending data to a slice may require allocating the data to a new array to accommodate the memory.

The built-in `append` function allows you to add data to a slice.

```go
hobbits := []string{ "Frodo", "Bilbo", "Samwise" }

append(hobbits, "Pippin")
```

## Two-dimensional Slices

2D slices (array of arrays, etc) in Go can be achieved with the `make` function or composite literal syntax.

With `new` function

```go
type TicTacToe [3][3]string

board := new(TicTacToe) // creates new TicTacToe 2d array with nil'd values
```

or with composite literal syntax

```go
type TicTacToe [3][3]string

board := TicTacToe{
    [3]string{"_", "_", "_"},
    [3]string{"_", "_", "_"},
    [3]string{"_", "_", "_"}
}
```

## Maps

In Go, _maps_ are data structures similar to _dictionary_ objects in other languages. Structures are defined with _key/value_ pairs, where the keys can be any data type that supports equality.

Use bracket notation to access a value at a given property. If a value does not exist, the zeroed value is returned.

````go
beastiary := map[string][]string {
    "hobbits": []string{ "Frodo", "Samwise", "Pippin", "Merry" },
    "humans": []string{ "Aragorn", "Faramir", "Boromir", "Theoden" },
    "dwarfs": []string{ "Gimli", "Thorin" }
}

beastiary["dwarfs"][0] // "Gimli"

Use a second returned value, `ok bool`, to check if a value exists on a map.

```go
hobbit := map[string]string {
    "name": "Frodo"
}

name, ok := hobbit["species"] // "", false
````

Use the built-in `delete` function to remove values at a given key on a map.

```go
delete(beastiary, "humans") // Deletes the value a key: "humans"
```

## Printing

Go has many methods for printing and formatting values accessible through the `fmt` package.

Most common are `fmt.Println()` and `fmt.Print()`, but `fmt.Printf()` can be used to format _and_ print the given string.

```go
fmt.Println("Hello", "World") // "Hello World\n"
fmt.Print("Hello", "World") // "HelloWorld"
fmt.Printf("Hello World %d", 123) // "Hello World 123"
```

Read more about the [fmt package](https://pkg.go.dev/fmt)

## Append

Go has a built-in `append` function for adding elements to the end of a slice. The first parameter is the slice with type, and the next arguments of arbitrary amount, are the items to append. It will return a slice.

Use the `...` operator to spread the values from one slice into another.

```go
hobbits := []string{"Bilbo", "Frodo"}
hobbits = append(hobbits, "Pippin", "Merry")

moreHobbits := []string{"Samwise", "Gaffer"}

hobbits = append(hobbits, moreHobbits...)
```
