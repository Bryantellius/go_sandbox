package a

import (
	"fmt"

	b "github.com/Bryantellius/go_sandbox/b"
)

func DoSomething() {
	fmt.Println("Beginning A...")

	fmt.Printf("A for addition. 1 + 1 = %d\n", Sum(1, 1))

	b.DoSomething()
}
