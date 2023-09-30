# Data Structure in Go

## Arrays

An array is a _fixed size_ collection of items of specified type.

Arrays are declared with `[n]T` syntax where `n` is the fixed length and `T` is the specified type of elements to be contained within the array.

> By default, array values will be _zeroed values_ of the specified type

```go
var scores [5]int // [0, 0, 0, 0, 0]
```

You can set array values based on _indices_. Use `arr[index]= value` syntax to assigned array values.

> Arrays are _zero-indexed_ meaning the first value is at index 0

```go
scores[0] = 1
```

You can also instantiate arrays with values with `[n]T{}` syntax.

```go
var teamMembers = [5]string{"Ben", "Hayden", "BG", "Hunter", "Jon"}
```

Lastly, arrays are one-dimensional data structures, but can hold array values, meaning you can represent multi-dimensional data.

```go
ticTacToeBoard := [3][3]string{ [3]{ "X", "O", "O" }, [3]{ "X", "O", "X" }, [3]{ "O", "X", "X" } }
```

## Slices

A slice is a _dynamically size_ collection of items of specified type. _It is a subsection of an underlying array._ Slices can be defined from existing arrays using `array[start:end]` syntax, or as new collections with `[]string{...}` syntax.

Empty slices can be created with the `make` function, with a length and capacity value being optionally supplied.

```go
var hobbits = [4]string{"Frodo", "Sam", "Merry", "Pippin"}

// declaring a slice from an array
var hobbitsWithRing = hobbits[1:3]
var hobbitsWithoutRing = hobbits[3:]

// declaring an empty slice (not from an existing array)
var extraCharacters = make([]string, 3) // creates a slice of length 3, capacity 3

// declaring a slice with values (not from an existing array)
var chasePack = []string{"Aragorn", "Legolas", "Gimli"}
```

## Maps

## `len()`

The `len()` function is an easy way to find the length of an array, slice, or map. Pass the data structure to the function and it will return the length.
