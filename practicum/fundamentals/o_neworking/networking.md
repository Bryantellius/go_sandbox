# Networking in Go

Go has a builtin standard package called `net/http` that enables network access from a go program.

## Retrieving Data from Network Requests

The `net/http` package exports functions that allow you to make network requests. Below is an example of calling a "GET" request.

```go
const url = "https://some.url.com"

func CallApi() {
    response, err := http.Get(url)

    checkError(err)

    defer response.Body.Close()

    bytes, err := io.ReadAll(response.Body)

    fmt.Println(strings(bytes))
}

```

## Handling JSON Data

The `encoding/json` package exports functions that allow you to encode, decode, and manage json data in Go.

We can call parse json content as follows:

```go
var terms []Term

// json.Unmarshal decodes json encoded data (slice of bytes)
// and stores the value in the given value
err = json.Unmarshal(bytes, &terms)
checkError(err)

for _, term := range terms {
    fmt.Println(term.Name)
}
```
