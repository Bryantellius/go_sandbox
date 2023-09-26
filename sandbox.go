package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

func main() {
	log.SetPrefix("main: ")
	log.SetFlags(0)

	names := []string{"Frodo", "Bilbo", "Sam"}
	messages, err := CreateMessages(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	for _, message := range messages {
		color.Green(message)
	}
}
