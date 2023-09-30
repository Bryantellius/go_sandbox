package main

import "fmt"

const LESSON_NUMBER int = 3

func PrintWhatILearned() {
	wallet := 14.00
	price := 12.99
	tax := .10

	if total := price + price*tax; wallet > total {
		fmt.Printf("Can afford. Wallet balance is now %f", wallet-total)
	} else {
		fmt.Printf("Cannot afford. Balance would need %f", total-wallet)
	}

	birthMonth := 1
	var monthStr string

	switch {
	case birthMonth == 12 || birthMonth < 3:
		monthStr = "Winter"
	case birthMonth < 6:
		monthStr = "Spring"
	case birthMonth < 9:
		monthStr = "Summer"
	case birthMonth < 12:
		monthStr = "Fall"
	}

	fmt.Println()
	fmt.Printf("You were born in %s", monthStr)
}

func main() {
	fmt.Printf("\n%d. Control Structures.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
