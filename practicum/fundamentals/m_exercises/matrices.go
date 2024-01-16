package main

import (
	"fmt"
	"sort"
)

type position struct {
	row int
	col int
}

// Traverses the values in a matrix by column, alternating direction at each column end.
func Traverse(mtx [][]int) [][]int {
	height := len(mtx)
	length := len(mtx[0])
	positions := make([]position, 0)
	values := make([]int, 0)

	dir := 1
	row := -1

	for col := 0; col < length; col++ {
		for {
			row += dir
			if dir > 0 && row < height || dir < 0 && row >= 0 {
				positions = append(positions, position{row, col})
				values = append(values, mtx[row][col])
			} else {
				break
			}
		}

		dir = -dir
	}

	// sort values
	sort.Ints(values)

	fmt.Println(values)
	fmt.Println(positions)

	// place sorted values in correct positions
	result := make([][]int, height)

	for idx := range result {
		result[idx] = make([]int, length)
	}

	for idx, pos := range positions {
		result[pos.row][pos.col] = values[idx]
	}

	return result
}

func main() {
	input := [][]int{{1, -1, 4, 1}, {7, -20, 12, 0}, {8, 10, -4, -3}}
	output := Traverse(input)

	fmt.Println(input)
	fmt.Println(output)
}
