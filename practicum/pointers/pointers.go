package main

import (
	"fmt"
	"strings"
)

const LESSON_NUMBER int = 6

func ToUpper(name *string) {
	*name = strings.ToUpper(*name)
}

func PrintWhatILearned() {
	var name string = "Ben"
	fmt.Println(name)
	ToUpper(&name)
	fmt.Println(name)
}

func main() {
	fmt.Printf("\n%d. Pointers.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
