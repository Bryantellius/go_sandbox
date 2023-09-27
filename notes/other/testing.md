# Testing

Unit tests can be created easily with Go's built-in `testing` library.

Create test files as `name_test.go`, then run the `go test` command.

Test functions should start with `Test` and receive a pointer param to the `testing.T` type from the `testing` built-in library:

```go
func TestHelloWorld(t *testing.T) {
    // ...test
}
```

A test passes if no `T` type method is called to end execution.

For example, let's complete `TestHelloWorld` that tests a `HelloWorld` function that should return the text `"Hello, World!"`.

```go
func TestHelloWorld(t *testing.T) {
    message := HelloWorld()
    want := "Hello, World!"

    if message != want {
        t.Fatalf(`HelloWorld() = "%s", want match for "%s"`, message, want)
    }
}
```

Here, we call the `HelloWorld` function, and have a control structure that determines if the received value matches our expectation. If not, we call `t.Fatalf` to pass a formatted message to be printed to the standardout.
