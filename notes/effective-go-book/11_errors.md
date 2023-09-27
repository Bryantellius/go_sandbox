# Effective Go: Errors

Go allows for multiple return values, making it appropriate and conventional to return an error value from a function.

Errors are structured as:

```go
type error interface {
    Error() string
}

// Specific Error
// PathError records an error and the operation and
// file path that caused it.
type PathError struct {
    Op string    // "open", "unlink", etc.
    Path string  // The associated file.
    Err error    // Returned by the system call.
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```

## Table of Contents

- [Panic](#panic)
- [Recover](#recover)

## Panic

Go provides a builtin `panic` method that creates a runtime error and terminates the program. The panic method can take any arbitrary value, but it is usually a string. 

Generally, it is best to work around or handle errors and continue program execution, rather than crashing the program.

## Recover
