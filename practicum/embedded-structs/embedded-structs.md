# Embedded Structs in Go

Go allows for _embedded structs and interfaces_ that offers better composition of types.

Embedded structs can be listed as fields when defining struct types.

```go
type time struct {
	hours   int
	minutes int
	seconds int
}


type runner struct {
	time // embedded struct
	fname string
	lname string
}
```

Embedded struct values can be access through dot-notation either top level or through the embedded struct field. The same goes for accessing embedded struct methods.

```go
r := runner{
    fname: "John",
    lname: "Doe",
    time: time{
        hours: 0,
        minutes: 16,
        seconds: 06,
    },
}

r.time.hours // 0
r.hours // 0
```

If embedded structs implement interfaces, then their super struct also implements the interface.
