package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello from Go!"
	file, err := os.Create("./data/hello_world.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %d characters\n", length)
	defer file.Close() // important to close files when finished
	defer readFile("./data/hello_world.txt")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:", string(data))
}
