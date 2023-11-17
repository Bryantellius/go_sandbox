# Loops in Go

The `for` loop is the only loop construct in Go.

It is made of the `for` keyword and three parts separated by semicolons:

- initializer (typically shorthand variable declaration)
- conditional
- post statement (usually an increment)

Loops will continue to execute their loop body until their condition equates to `false`.

```go
for i := 1; i <= 10; i++ {
    fmt.Printf("i is %d", i)
    fmt.Println()
}

// i is 1
// i is 2
// i is 3
// i is 4
// i is 5
// i is 6
// i is 7
// i is 8
// i is 9
// i is 10
```

## Variations

The initializer and post action statement are optional. If these parts are omitted, you still need to separate their sections with semicolons.

```go
i := 1

for ; i <= 10; i+=2 {
    fmt.Printf("i is %d", i)
    fmt.Println()
}

// i is 1
// i is 3
// i is 5
// i is 7
// i is 9
```

or

```go
for i := 1; i <= 10; {
    fmt.Printf("i is %d", i)
    fmt.Println()
    i += 5
}

// i is 1
// i is 6
```

If both the initializer and action statement are omitted, you can drop the semicolons altogether (similar to `while` loop syntax in other languages).

```go
result := 1
multiplier := 10

for result <= 1000 {
    fmt.Printf("result is %d", result)
    fmt.Println()
    result *= multiplier
}

// result is 1
// result is 10
// result is 100
// result is 1000
```

## Additional Keywords

There are two keywords that work specifically with loops.

- `break` stops the loop
- `continue` stops the iteration and continues at the next iteration

```go
numToFind := 7
num := 0

for {
    num++

    if num > 0 && num%2 == 0 {
        fmt.Printf("Skipping %d\n", num)
        continue
    } else if num == numToFind {
        break
    }

    fmt.Printf("Stepping over %d\n", num)
}

fmt.Printf("Found %d", numToFind)
```
