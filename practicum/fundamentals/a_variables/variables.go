package main

import (
	"fmt"
	"math"
)

const LESSON_NUMBER int = 1

func PrintWhatILearned() {
	// declaring a variable with explicit type and assigning a value
	var myName string = "Ben Bryant"

	// declaring a multiple variables with same explicit type and assigning values
	var month, day, year int = 7, 1, 1997

	// declaring multiple variables with different explicit types and assigning values (factored declarations)
	var (
		username   string = "Bryantellius"
		isActive   bool   = true
		avatarChar rune   = 'B'
	)

	// declaring and assigning a variable in shorthand with inferred typing
	dob := fmt.Sprintf("%d/%d/%d", month, day, year)

	fmt.Println()
	fmt.Printf("Hello! My name is %s and I was born on %s. My profile is \n- Active: %v\n- Username: %s\n- Avatar: %c", myName, dob, isActive, username, avatarChar)
	fmt.Println()
	fmt.Printf("The largest number that can be assigned as a int is %d", math.MaxInt64)
}

func main() {
	fmt.Printf("\n%d. Declaring, assigning and initializing variables.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
