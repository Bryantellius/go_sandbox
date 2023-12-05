package main

import (
	"fmt"
)

const LESSON_NUMBER int = 7

type runner struct {
	name  string
	place int
}

func createRunner(name string) *runner {
	newRunner := runner{name: name}
	return &newRunner
}

func PrintWhatILearned() {
	bb := createRunner("Ben")
	hw := createRunner("Hayden")

	fmt.Println(bb)
	fmt.Println(hw)

	bb.place = 1
	hw.place = 13

	fmt.Println(bb)
	fmt.Println(hw)

	// anonymous struct type
	value := struct {
		mode     string
		username string
	}{
		"Easy",
		"Bryantellius",
	}

	fmt.Println(value)
}

func main() {
	fmt.Printf("\n%d. Structs.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
