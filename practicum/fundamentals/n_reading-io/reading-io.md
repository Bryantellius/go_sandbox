# Reading Input Output in Go

## Local File System

Go has builtin packages that enable access to input output devices and the operating system. They are aptly named `io` and `os` respectively.

### Reading and Writing Local Files

To write local files, we can use the `os.Create` and `io.WriteString` functions as shown below:

```go
// os.Create takes in a file name
// returns a os.File struct representing the file
// or an error if one occurred
file, err := os.Create("./path/to/file")

// io.WriteString takes in a struct that implements io.Writer
// and a string to write to the file
// returns the length of characters written
// or an error if one occurred
length, err := io.WriteString(file, "Some content...")

// after you finish interacting with a file,
// it is imperative to close the file
file.Close()
```

Similarly, to read local files, we can use the `os.ReadFile` function as shown below:

```go
// os.ReadFile takes in a string file name
// returns a slice of bytes
// or an error if one occurred
data, err := os.ReadFile("./path/to/file")

// the slice of bytes can be converted to a string for human readability
fmt.Println(string(data))
```
