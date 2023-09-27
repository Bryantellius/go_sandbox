# Effective Go: Syntax

## Table of Contents

- [Formatting](#formatting)
- [Commentary](#commentary)
- [Names](#names)
  - [Package names](#package-names)
  - [Getters](#getters)
  - [Interface names](#interface-names)
  - [MixedCaps](#mixedcaps)
- [Semicolons](#semicolons)

## Formatting

Use `go fmt` to format your Go program to a prescribed style.

_Indentions_ are denoted by **tabs**.

There is no _line length_ limit, so if a line is too long, wrap the line and _indent_.

There are fewer parentheses used in Go.

Operator precedence is denoted by _spaces_.\*

\* May be in Go 1.12 or prior. Not sure about Go 1.18.

## Commentary

Comments are denoted by:

- `//` for single line
- `/* */` for multi-line

`godoc` is used to produce documentation from top-level comments in a package.

Every package needs to have a _multi-line_ package comment that precedes the package declaration that will be used by `godoc` to produce commentary about a package. The commentary will be plain text and formatted similar to using `go fmt`. Do not include stars, banners, etc.

```go
/*
Package somepackage implements ...
*/
package somepackage
```

Every _exported name_ (starting with a capitalized letter) should have a _doc comment_. A doc comment starts with the declared name, and contains a sentence summary of the value.

```go
// Sum recieves two integers and returns the sum.
func Sum(x int, y int) int {
    res := x + y
    return res
}
```

Go allows for _grouped declarations_, and comments before grouped declarations can be helpful.

```go
// Form values used in server-side validation
var (
    age int
    dob string
    name string
)
```

## Names

Generally, _exported names_ are written with _pascal casing_.

### Package Names

Package names are concise, _lower case_ names that are accessors to exported names in a package.

Generally, there is no collision with names, since package names can use aliases per file, and often package names use the names of their source directories.

```go
import "fmt"

func main() {
    fmt.Println("Using the 'fmt' package to access 'Println'")
}
```

### Getters

Getters, if set, are generally given the field name exported in _upper case_. Setters by convention are named `Set<Field>`.

```go
// Getter for hobbit.name
hobbit.Name()
// Setter for hobbit.name
hobbit.SetName()
```

### Interface Names

One method interfaces are named by the method name plus "-er".

```go
type Stringer interface {
    String() string
}
```

### Mixed Caps

Go naming convention revolves around _mixed caps_. Use _camelCase_ for unexported names. Use _PascalCase_ for exported names.

## Semicolons

Semicolons are automatically inserted before newline characters. The only necessary places for semicolons are between initialization, condition and action of for loops.
