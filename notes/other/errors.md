# Errors

Go Errors are typically handled by returning multiple values from a function (the error being the second value). The function can determine if an error occurs and through it's control structure, create a new error and return it. If an error does not occur, return `nil`.

For example, let's say we want to create a function that handles integer division:

```go
func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	} else {
		result := a / b
		return result, nil
	}
}
```

`Divide` takes in two integer parameters, and returns the result integer _and an error_. Now, if there is no error, like after dividing `10 and 5`, then the function would yield `2, nil`. But if the function was given `10 and 0`, then the function would yield `0, error`.

In this way, the calling code can execute this function and check to see if an error was returned. 