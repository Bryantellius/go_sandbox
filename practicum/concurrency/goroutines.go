package main

import (
	"fmt"
	"time"
)

const LESSON_NUMBER int = 12

func printAllMultiplesTo100(factor int) {
	for i := 0; i <= 100; i += factor {
		time.Sleep(5 * time.Millisecond)
		fmt.Printf("%d: %d\n", factor, i)
	}
}

func PrintWhatILearned() {
	go printAllMultiplesTo100(5)
	go printAllMultiplesTo100(10)
	printAllMultiplesTo100(3)
}

func main() {
	fmt.Printf("\n%d. Goroutines.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
