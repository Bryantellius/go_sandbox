package b

import (
	"fmt"
	"strings"
)

func DoSomething() {
	fmt.Println("Beginning B...")

	beCool()
}

func beCool() {
	coolLetters := []string{"6", "3", "c", "0", "0", "1"}
	fmt.Printf("%s\n", strings.Join(coolLetters, " "))
}
