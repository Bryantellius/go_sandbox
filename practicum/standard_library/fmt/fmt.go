package main

import (
	"bufio"
	"fmt"
	"os"
)

type circle struct {
	radius, border int
}

func main() {
	// printing to the console

	fmt.Print("Welcome to Go!")

	fmt.Println("This ends the string with a newline.")

	var myName string = "Ben"

	fmt.Println("My name is", myName)

	scores := []int{1, 2, 3, 4, 5}
	length, err := fmt.Println(scores)
	fmt.Println(length, err)

	// formatting and printing with Printf

	x := 20
	f := 123.4567

	fmt.Printf("%d\n", x)
	fmt.Printf("%x\n", x)

	fmt.Printf("%t\n", x > 10)

	fmt.Printf("%f\n", f)
	fmt.Printf("%e\n", f)

	fmt.Printf("%d %#[1]o %#[1]x\n", 52)

	c := circle{20, 5}

	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%T\n", c)

	// formatting with Sprintf

	output := fmt.Sprintf("%d %#[1]o %#[1]x", 52)
	fmt.Println(output)

	// formatting with precision

	fmt.Printf("%.2f\n", f)
	fmt.Printf("%10f\n", f)
	fmt.Printf("%10.2f\n", f)
	fmt.Printf("%+10.2f\n", f)
	fmt.Printf("%010.2f\n", f)

	// reading from standard input

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}
