package utils

import (
	"regexp"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	name := "Frodo"
	message, err := CreateMessage(name)
	want := regexp.MustCompile(`\\s` + name + `$`)

	if want.MatchString(message) || err != nil {
		t.Fatalf(`CreateMessage("Frodo") = %s, %v, want match for %s, nil`, message, err, want)
	}
}
