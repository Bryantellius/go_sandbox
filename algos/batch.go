package algos

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Write a program which reads fictitious commands from a file.
// The commands should be in the form "operator operand".
// Print out a log of what the commands are in the format "Executing operator on operand"

func RunBatch(filepath string) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, fileInfo.Size())
	file.Read(data)

	if err != nil {
		log.Fatal(err)
	}

	parsedCommands := parseFileForCommands(data)

	printCommands(parsedCommands)
}

// Converts the file data from a slice of bytes to command strings.
// The final format will resemble [ ["command", "value"], ... ]
func parseFileForCommands(fileData []byte) [][]string {
	fileDataStr := string(fileData)
	fileDataLines := strings.Split(fileDataStr, "\r\n")

	result := make([][]string, len(fileDataLines))

	for idx, line := range fileDataLines {
		result[idx] = strings.Split(line, " ")
	}

	return result
}

// Prints the values from the parsed commands
func printCommands(data [][]string) {
	start := time.Now()
	fmt.Printf("Starting Execution\t%s\n=======================\n", start)
	for _, line := range data {
		fmt.Printf("Executing %s on %s\t%s\n", line[0], line[1], time.Now())
	}
	fmt.Printf("=======================\nFinished Execution\t%s", time.Now().Sub(start))
}