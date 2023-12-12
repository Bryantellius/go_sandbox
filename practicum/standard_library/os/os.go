package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	exists := checkIfFileExists("./os.md")
	fmt.Println("File exists:", exists)

	stats, err := os.Stat("./os.md")
	if err != nil {
		panic(err)
	}

	fmt.Println("Modification time:", stats.ModTime())
	fmt.Println("Modification Mode:", stats.Mode())
	fmt.Println("File Size:", stats.Size())
	fmt.Println("Is directory:", stats.IsDir())
}

func checkIfFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func createFile(data []byte, filePath string) {
os.WriteFile(filePath, data, os.)
}