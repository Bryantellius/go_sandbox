# Compiling and Installing Your Application

You can build your project with the `go build` command.

This will compile the project and generate a _binary executable file_.

This executable can be run from the terminal locally if you are in the right folder.

You can also install the executable to a specific location by:

1. Discovering the install path with `go list -f '{{.Target}}'`
2. Adding the Go install directory to the shell path
   - Linux/Unix `export PATH=$PATH:/path/to/dir`
   - Windows `set PATH=%PATH%;/path/to/dir`
3. Run `go install`

Then you will be able to execute the application from any dir by using the name of the application.
