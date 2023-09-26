# Working with Modules

## Creating a Module

Go code is grouped into packages that are grouped into modules. Each module specifies the version and dependencies to run the code.

To create a Go module, run:

- `go mod init <domain>/<module>`

For example, if you have your remote destination as Github, and your module (root working directory name) is "mymodule", then you would run:

- `go mod init github.com/mymodule`

`go mod init` creates a `go.mod` file to keep track of the name, version, and dependencies of your module.

## Calling code from a Module

Calling code from another module (typically an external remote module) can be done by:

- searching/finding the module/package at [https://pkg.go.dev/](https://pkg.go.dev/)
- adding an `import "domain/package"` statement to your file
- running `go mod tidy` for go to fetch and authenticate the required package
