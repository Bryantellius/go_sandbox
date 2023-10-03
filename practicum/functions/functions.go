package main

import (
	"errors"
	"fmt"
)

const LESSON_NUMBER int = 5

func sum(values ...int) int {
	result := 0

	for _, value := range values {
		result += value
	}

	return result
}

func IsNotZero(n int) (bool, error) {
	if n != 0 {
		return true, nil
	} else {
		return false, errors.New("Input cannot be 0")
	}
}

func mapArr(arr []int, fn func(int) int) []int {
	var resultArr []int = make([]int, len(arr))

	for index, value := range arr {
		resultArr[index] = fn(value)
	}

	return resultArr
}

func newId() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

func isEven(n int) bool {
	if n == 1 {
		return false
	} else if n == 0 {
		return true
	} else {
		return isEven(n - 2)
	}
}

func PrintWhatILearned() {
	var numbers []int = []int{1, 2, 3, 4, 5}
	fmt.Println(sum(numbers...))
	fmt.Println(mapArr(numbers, func(n int) int { return n * 2 }))
	fmt.Println(isEven(10))

	generateId := newId()

	fmt.Println(generateId())
	fmt.Println(generateId())
	fmt.Println(generateId())
}

func main() {
	fmt.Printf("\n%d. Functions.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
