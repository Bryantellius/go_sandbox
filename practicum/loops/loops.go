package main

import "fmt"

const LESSON_NUMBER int = 2

func PrintWhatILearned() {
	numToFind := 7
	num := 0

	for {
		num++

		if num > 0 && num%2 == 0 {
			fmt.Printf("Skipping %d\n", num)
			continue
		} else if num == numToFind {
			break
		}

		fmt.Printf("Stepping over %d\n", num)
	}

	fmt.Printf("Found %d", numToFind)
}

func main() {
	fmt.Printf("\n%d. For loop.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
