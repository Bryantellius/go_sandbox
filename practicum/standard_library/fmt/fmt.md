# The `fmt` Standard Package

## Printing

There are two main functions in the `fmt` package for printing to the standard output: `Print` and `Println`.

### `Print`

The `Print` function is a variadic function and prints the arguments to the standard output.

```go
// signature
func Print(a ...any) (n int, err error)

// example
fmt.Print("Welcome to Go!")
```

### `Println`

The `Print` function is a variadic function and prints the arguments to the standard output, ending in a carriage return.

```go
// signature
func Println(a ...any) (n int, err error)

// example
fmt.Println("This ends the string with a newline.")
```

## Formatting

There are two main functions in the `fmt` package for formatting string outputs: `Printf` and `Sprintf`. There are many _verbs_ or _flags_ used to control the formatting. Below are some common formatting verbs in Go:

| Verb | Description           |
| ---- | --------------------- |
| `%d` | base 10 integer       |
| `%f` | floating point number |
| `%t` | boolean               |
| `%v` | default value format  |
| `%T` | value type            |
| `%s` | string                |

### `Printf`

The `Printf` function is a variadic function and combines the arguments given to the format string, and prints the arguments to the standard output.

```go
// signature
func Printf(format string, a ...any) (n int, err error)

// example
name := "Frodo"
age := 51
fmt.Printf("%s was %d when he started his journey.", name, age)
```

### `Sprintf`

The `Sprintf` function is a variadic function and combines the arguments given to the format string, and returns the string output.

```go
// signature
func Sprintf(format string, a ...any) string

// example
name := "Frodo"
age := 51
output := fmt.Sprintf("%s was %d when he started his journey.", name, age)
```

## Reading Input from Standard Input

One way to read input from standard input in Goland is to use functions from the `bufio` and `os` packages.

You can use the `bufio.NewReader` function and pass in the `os.Stdin` variable to create a buffered I/O reader based on the OS's standard input. From there, you can call the `io.Reader`'s `ReadString` method to read strings that are inputted via the standard input, and then use those values in your program.

> The `io.Reader` has other methods that can allow you to read more than just string values.

```go
reader := bufio.NewReader(os.Stdin)

fmt.Println("Enter your name: ")

s, err := reader.ReadString('\n')

if err != nil {
    panic(err)
}

fmt.Println("Hello", s)
```
