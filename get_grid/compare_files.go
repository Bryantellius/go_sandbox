package getgrid

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

func CompareCSVsAndRemoveDuplicates(filenameA, filenameB, outputFilename string) {
	contentsA := ReadCSVFile(filenameA)
	contentsB := ReadCSVFile(filenameB)

	headersA, recordsA := ParseCSV(contentsA)
	headersB, recordsB := ParseCSV(contentsB)

	if !slices.Equal(headersA, headersB) {
		log.Println("WARNING: Headers are not the same for compared csv files.")
	}

	results := [][]string{headersA}
	var longer [][]string
	var shorter [][]string

	if len(recordsA) >= len(recordsB) {
		longer = recordsA
		shorter = recordsB
	} else {
		shorter = recordsA
		longer = recordsB
	}

	for _, record := range longer {
		isFound := false
		for _, r := range shorter {
			if slices.Equal(record, r) {
				isFound = true
				break
			}
		}
		if !isFound {
			results = append(results, record)
		}
	}

	for _, record := range shorter {
		isFound := false
		for _, r := range longer {
			if slices.Equal(record, r) {
				isFound = true
				break
			}
		}
		if !isFound {
			fmt.Println(record)
		}
	}

	log.Printf("INFO: Found %d unique records\n", len(results)-1)
	// log.Printf("INFO Results:\n%v", results)
	WriteCSVFile(outputFilename, results)
	log.Printf("INFO Results written to file: %s", outputFilename)
}

func ParseCSV(contents []byte) ([]string, [][]string) {
	rA := csv.NewReader(bytes.NewReader(contents))

	var headers []string
	var records [][]string

	for {
		record, err := rA.Read()
		if err == io.EOF {
			break
		}
		handleErr(err)

		if len(headers) == 0 {
			headers = append(headers, record...)
		} else {
			records = append(records, record)
		}
	}

	return headers, records
}

func ReadCSVFile(filename string) (contents []byte) {
	exists := checkIfFileExists(filename)

	if !exists {
		panic("File does not exist.")
	} else {
		buffer, err := os.ReadFile(filename)
		handleErr(err)
		return buffer
	}
}

func WriteCSVFile(filename string, contents [][]string) {
	file, err := os.Create(filename)
	handleErr(err)

	defer file.Close()

	w := csv.NewWriter(file)
	err = w.WriteAll(contents)
	handleErr(err)
}

func checkIfFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
