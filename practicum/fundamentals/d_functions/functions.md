# Functions in Go

Functions are central to Go programming and allow you to reuse code blocks with same of differing inputs to complete your program.

The `func` keyword is used to declare a function, followed by the name of a function. Function naming conventions follow the same rules as variables names.

Parameters must be statically typed the same as variables.

If a function returns a value, the return data type must be explicitly set. Any value following the `return` keyword is returned from a function.

```go
func SumArr(values []int) int {
    result := 0

    for _, value := range values {
        result += value
    }

    return result
}
```

## First Class Citizens

Functions are first class citizens in Go and can be assigned to variables, passed as parameters, assigned to key values, and returned from functions.

```go
func mapArr(arr []int, fn func(int) int) []int {
    var resultArr []int = make([]int, len(arr))

    for index, value := range arr {
        resultArr[index] = fn(value)
    }

    return resultArr
}
```

## Named Return Values

Return values can be named. In shorter functions, named values can enable _naked returns_ where the returned variables are implicitly returned instead of explicitly listed in the return statement.

## Multiple Return Values

Functions can have multiple return values. A common practice in _idiomatic go_ is returning a value and error from a function to handle propagated errors.

```go
func IsNotZero(n int) (bool, error) {
    if n != 0 {
        return true, nil
    } else {
        return false, errors.New("Input cannot be 0")
    }
}

// ...

result, err := IsNotZero(0)
```

## Variadic Functions

Functions can be created to take an arbitrary number of parameters that are gathered in a slice. Use `...` to gather parameters of a specified type, or provide a slice followed by `...` to spread the values of a slice into a variadic function.

```go
func sum(values ...int) int {
    result := 0

    for _, value := range values {
        result += value
    }

    return result
}

// ...
var numbers []int = []int{1, 2, 3, 4, 5}
fmt.Println(sum(numbers...))
```

## Closure

When a function references an outside variable and is returned from another function, that variable is wrapped in a `closure` where the function can reference and update the variable value outside of its defined scope.

```go
func newId() func() int {
    id := 0
    return func() int {
        id++
        return id
    }
}
```

## Recursion

A function can call itself, resulting in _recursion_. This results in functionality similar to loops.

> Be wary that if there is no condition or flow that allows the recursive function call to end, it will result in an infinite loop.

```go
func isEven(n int) bool {
    if n == 1 {
        return false
    } else if n == 0 {
        return true
    } else {
        return isEven(n - 2)
    }
}
```
