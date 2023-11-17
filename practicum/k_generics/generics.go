package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

const LESSON_NUMBER int = 11

// Generic Type Set Interface
type CommonNumber interface {
	int | float64
}

// type with generic parameter types
type Person[T constraints.Ordered] struct {
	id   T
	name string
}

// function with generic parameter types
func add[T CommonNumber](x, y T) T {
	return x + y
}

func PrintWhatILearned() {
	// with explicit type argument during instantiation
	resultInt := add[int](1, 2)

	fmt.Println(resultInt)

	// without explicit type argument during instantiation
	// with function argument type inference
	resultFloat := add(3.15, 3.85)

	fmt.Println(resultFloat)

	frodo := Person[int]{1, "Frodo"}

	fmt.Println(frodo)

	bilbo := Person[string]{"123b", "Bilbo"}

	fmt.Println(bilbo)
}

func main() {
	fmt.Printf("\n%d. Generics.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
