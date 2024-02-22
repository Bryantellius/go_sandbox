package utils

import (
	"errors"
	"fmt"
	"math/rand"
)

var formats = []string{
	"Hello %v",
	"Hi %v",
	"Welcome %v",
}

func CreateMessage(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	} else {
		var selectedFormat string = formats[rand.Intn(len(formats))]
		var message string = fmt.Sprintf(selectedFormat, name)
		return message, nil
	}
}

func CreateMessages(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for index, name := range names {
		message, err := CreateMessage(name)

		fmt.Printf("Creating message %d for %s\n", index, name)

		if err != nil {
			return nil, err
		} else {
			messages[name] = message
		}
	}

	return messages, nil
}
