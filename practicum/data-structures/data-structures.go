package main

import "fmt"

const LESSON_NUMBER int = 4

func Arrays() {
	// declare an array with zeroed values
	var scores [5]int

	scores[0] = 1
	scores[1] = 2
	scores[2] = 3
	scores[3] = 4
	scores[4] = 5

	fmt.Println(scores[2])

	// declare an array with set values
	var teamMembers = [5]string{"Ben", "Hayden", "BG", "Hunter", "Jon"}

	fmt.Println(teamMembers[4])

	// declare an array shorthand with set values of array types
	ticTacToeBoard := [3][3]string{{"X", "O", "O"}, {"X", "O", "X"}, {"O", "X", "X"}}

	fmt.Println(ticTacToeBoard)
}

func Slices() {
	var hobbits = [4]string{"Frodo", "Sam", "Merry", "Pippin"}

	// declaring a slice from an array
	var hobbitsWithRing = hobbits[0:2]
	var hobbitsWithoutRing = hobbits[2:]

	// declaring an empty slice (not from an existing array)
	var extraCharacters = make([]string, 3) // creates a slice of length 3, capacity 3

	// declaring a slice with values (not from an existing array)
	var chasePack = []string{"Aragorn", "Legolas", "Gimli"}

	fmt.Println(hobbitsWithRing)
	fmt.Println(hobbitsWithoutRing)
	fmt.Println(extraCharacters)
	fmt.Println(chasePack)
}

func PrintWhatILearned() {
	Arrays()

	Slices()
}

func main() {
	fmt.Printf("\n%d. Data Structures.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
