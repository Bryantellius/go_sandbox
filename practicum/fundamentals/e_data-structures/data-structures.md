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

Slices can also contain slice values to create _jagged arrays_ (collections with collection values of different lengths).

### `append`

The `append()` function is used to add values to end of a slice. If the slice has enough _capacity_ for the value, then the slice is re-sliced from the underlying array and  accommodates the value. If not, then a new underlying array is created to hold the new elements.

### The `slices` Package

The `slices` package has many helpful functions for working with slices.

## Maps

`Maps` are data collections that map keys to values.

A map can be created with the `make()` function and providing the key/value types with `map[key_type]value_type` syntax.

```go
var teamScores map[string]int = make(map[string]int)
```

Map values can be accessed and assigned similar to arrays and slice syntax:

```go
teamScores["Ben"] = 1
teamScores["Kane"] = 6
teamScores["Kevin"] = 8
teamScores["BG"] = 12
teamScores["Andrew"] = 14
```

Maps can also be created as _map literals_ using the following syntax:

```go
fellowship := map[string]string{
    "Frodo":   "Hobbit",
    "Aragorn": "Man",
    "Gimli":   "Dwarf",
    "Legolas": "Elf",
}
```

Lastly, you can access a boolean value that specifies if a certain key exists on a map. When you access a key on a map, the zeroed value will be returned if the value has not be defined or if the key has not been defined. The second return value when accessing a map value is helpful in distinguishing the existence of keys.

```go
andrewsScore, ok := teamScores["Andrew"]
```

### `delete()`

The `delete()` builtin function is used to remove a key/value pair from a map.

```go
delete(teamScores, "Ben")
```

## `len()`

The `len()` function is an easy way to find the length of an array, slice, or map. Pass the data structure to the function and it will return the length.

## `copy()`

Arrays, slices and maps can be copied with the `copy()`

## `clear()`

The `clear()` builtin function is used to remove all key/value pairs from a map or slice.

```go
clear(fellowship)
```

## `range`

The `range` keyword allows you to iterate over keys and values in arrays, slices, and maps, as well as other data structures.

It is used in a `for` loop statement and returns two values, the key/index and the value.

> If you do not need the first return value of a function, you can use the `_` **blank identifier**.

```go
var totalScore int

for runner, score := range teamScores {
    fmt.Printf("%s scored %d", runner, score)
    totalScore += score
}

fmt.Printf("Total Team Score: %d", totalScore)
```
