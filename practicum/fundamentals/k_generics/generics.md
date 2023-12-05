# Generics in Go

Generics in Go allow you to create functions that take _generic type parameters_, create structs with _generic types_, and create interfaces with _generic type sets_.

## Type Parameters

Sometimes functions can be used with several different parameter and return types. For example, a function that _adds_ two numbers and returns the result can be written the same was for integers, floating point numbers, etc.

```go
func addInts(x, y int) int {
    return x + y
}

func addFloats(x, y float64) float64 {
    return x + y
}
```

Generic type parameters allow you to create one function that can take parameters of generic types.

```go
func add[T constraints.Ordered](x, y T) T {
	return x + y
}

// ...

resultInt := add[int](1, 2)

resultFloat := add[float64](3.15, 3.85)
```

Here is the add function with generic type parameters. Generic type parameters are listed in `[]` after the function name and before the parameter list.

Then to use the function, we have a two-step process. The function is _instantiated_ with the given type parameter, and then called with the given parameters.

## Generic Types

Generic parameter types can also be used with defining types.

```go
type Person[T constraints.Ordered] struct {
	id   T
	name string
}

// ...

frodo := Person[int]{1, "Frodo"}

bilbo := Person[string]{"123b", "Bilbo"}
```

They are _instantiated_ the same way as functions with providing the generic types in `[]` before the creation of the struct.

## Type Sets

Interfaces are defined as a collection of methods, and a type implements an interface by containing the defined methods. However, another way to think about it is that _interfaces are sets of types that implement them_. This way, interfaces can be used to create _generic type constraints_.

```go
// Generic Type Set Interface
type CommonNumber interface {
	int | float64
}

// ...

// function with generic parameter types
func add[T CommonNumber](x, y T) T {
	return x + y
}
```

## Type Inference

Earlier we explicitly defined argument type parameters when we called the `add` function. However, for most generic type parameters, they can be _inferred_ at compile time and we can omit the explicit argument types during instantiation of the function.

```go
// with explicit type argument during instantiation
resultInt := add[int](1, 2)

// without explicit type argument during instantiation
// with function argument type inference
resultFloat := add(3.15, 3.85)
```
