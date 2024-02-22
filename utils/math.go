package utils

import "errors"

func Sum(a int, b int) int {
	result := a + b
	return result
}

func Subtract(a int, b int) int {
	result := a - b
	return result
}

func Multiply(a int, b int) int {
	result := a * b
	return result
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	} else {
		result := a / b
		return result, nil
	}
}
