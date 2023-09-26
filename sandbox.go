package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

func main() {
	log.SetPrefix("main: ")
	log.SetFlags(0)

	message, err := CreateMessage("Ben")

	if err != nil {
		log.Fatal(err)
	}

	color.White(message)

	result, err := Divide(10, 0)

	if err != nil {
		log.Fatal(err)
	}

	color.White(fmt.Sprintf("Answer is: %d", result))
}
