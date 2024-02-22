package a

import (
	"fmt"
	utils "github.com/Bryantellius/go_sandbox/utils"
	b "github.com/Bryantellius/go_sandbox/b"
)

func DoSomething() {
	fmt.Println("Beginning A...")

	fmt.Println("A for addition.")

	addition()

	b.DoSomething()
}

func addition() {
	fmt.Printf("1 + 1 = %d\n", utils.Sum(1, 1))
}
