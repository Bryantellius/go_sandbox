# Getting Started Notes

When you create a new directory of Go code, you can enable dependency tracking by running:

- `go mod init domain/module`

Go files have the `.go` file extension.

Things within a file:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

`package main` groups functions within a directory
`import "fmt"` imports the fmt standard library
`func main` declares the main function that executes by default when you run a package

External packages can be found at [https://pkg.go.dev/](https://pkg.go.dev/).

You can use external packages by similarly importing the module/package at the top of a file:

```go
import "rsc.io/quote"

// ...

quote.Go()
```

When you use an external package for the first time, you need to run `go mod tidy` for Go to fetch the required module and authenticate it with a `go.sum` file.
