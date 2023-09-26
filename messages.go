package main

import (
	"errors"
	"fmt"
)

func CreateMessage(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	} else {
		var message string = fmt.Sprintf("Hello, %v", name)
		return message, nil
	}
}
