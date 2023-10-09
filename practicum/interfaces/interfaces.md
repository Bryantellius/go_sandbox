# Interfaces in Go

An _interface_ is a named collection of method signatures. Interfaces are inferred implicitly for any type that implements the interfaces methods. This allows any type in any collection to implement a collection.

```go
type Shape interface {
	Area() float64
}

type Square struct {
	width float64
}

func (s Square) Area() float64 { // implements Shape
	return s.width * s.width
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 { // implements Shape
	return r.width * r.height
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 { // implements Shape
	return t.base * t.height * 0.5
}

func printArea(s Shape) { // can receive any type that implements Shape
	fmt.Printf("%T Area: %f", s, s.Area())
}

//...

square := Square{10.0}
rectangle := Rectangle{10.0, 12.0}
triangle := Triangle{7.25, 11}

printArea(square)
fmt.Println()
printArea(rectangle)
fmt.Println()
printArea(triangle)
```

Interfaces are essentially _tuples_ that hold a value and a concrete type. Because of this, the value can be `nil` and the concrete type with receive a `nil` receiver. In other languages this would trigger a `null pointer exception`, but not in Go.

`nil` interface values will throw a runtime error since there are no concrete methods to call.

```go
func main() {
    var square Square // creates a square with a nil value; interface will still call concrete methods
    var s Shape // creates a nil interface value
}
```

_Empty interfaces_ are possible and useful when the type of the concrete value is unknown.

```go
var i interface{} // can be assigned any value

func takeAny(i interface{}) { // can accept any value type
    // ...
}
```
