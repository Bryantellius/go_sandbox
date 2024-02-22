package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput() {
	reader := *bufio.NewReader(os.Stdin)

	fmt.Print("Enter you name: ")

	strInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Hello ", strInput)
	}

	fmt.Print("Enter your current balance: ")
	numInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err.Error())
	} else {
		num, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("You entered: %.2f", num)
		}
	}
}
