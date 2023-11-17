package main

import (
	"fmt"
)

const LESSON_NUMBER int = 9

type MyError string

func (e MyError) Error() string { // MyError implements error interface with this method
	return fmt.Sprintf("BANG: %v", string(e))
}

type Shape interface {
	Area() (float64, error)
}

type Square struct {
	width float64
}

func (s Square) Area() (float64, error) {
	if s.width == 0 {
		return 0, MyError("must have width")
	}

	return s.width * s.width, nil
}

func (s Square) String() string { // implements the built-in Stringer interface
	return fmt.Sprintf("Square{%v}", s.width)
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() (float64, error) {
	if r.width == 0 || r.height == 0 {
		return 0, MyError("must have width and height")
	}

	return r.width * r.height, nil
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() (float64, error) {
	if t.base == 0 || t.height == 0 {
		return 0, MyError("must have base and height")
	}

	return t.base * t.height * 0.5, nil
}

func printArea(s Shape) {
	area, err := s.Area()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%T Area: %f", s, area)
	}
}

func PrintWhatILearned() {
	square := Square{0}
	rectangle := Rectangle{width: 10.0}
	triangle := Triangle{base: 7.25}

	printArea(square)
	fmt.Println()
	printArea(rectangle)
	fmt.Println()
	printArea(triangle)
	fmt.Println()
	fmt.Println(square)
}

func main() {
	fmt.Printf("\n%d. Interfaces.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
