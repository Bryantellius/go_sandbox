# Effective Go: The Blank Identifier

## Table of Contents

- [The blank identifier in multiple assignment](#the-blank-identifier-in-multiple-assignment)
- [Unused imports and variables](#unused-imports-and-variables)
- [Import for side effect](#import-for-side-effect)
- [Interface checks](#interface-checks)

## The Blank Identifier in Multiple Assignment

_Blank identifiers_ are `_` underscore symbols used to denote that a value is irrelevant. You can generally see this used with multiple return values where you only need to consider one of the two (or more).

```go
if _, err := os.Stat(path); os.IsNotExist(err) {
    fmt.Printf("%s does not exist\n", path)
}
```

## Unused Imports and Variables

Unused imports and variables with raise errors. If you want to silence these errors during development, use blank identifiers below the import statements with comments to delete later.

```go
import (
    "fmt"
    "sort"
)

var _ = sort.Sort // For debugging; delete when done

func Main() {
    fmt.Println("Example that uses fmt but not sort yet. The blank identifier above with allow this program to run")
}
```

## Import for Side Effects

There are some packages that provide side effects in their `Init` function. One example is `net/http/pprof` that provides debugging handlers for http requests. To import a package yet only use it's side effects, use a blank identifier:

```go
import _ "net/http/pprof"
```

## Interface Checks

You can verify that types satisfy interfaces with _type assertions_, but often you may only care about the returned `ok` value. Use a blank identifier in these cases.

```go
if _, ok := val.(json.Marshaler); ok {
    fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
}
```
