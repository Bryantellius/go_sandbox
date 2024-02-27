package main

import (
	"fmt"

	"github.com/Bryantellius/go_sandbox/algos"
)

func main() {
	DoCallStack()
}

func DoCallStack() {
	fmt.Println("Starting")
	callStack := algos.CreateStack()
	callStack.ShowStack()
	callStack.Push("main()")
	callStack.ShowStack()
	callStack.Push("DoSomething()")
	callStack.ShowStack()
	callStack.Pop()
	callStack.Push("DoSomethingReallyCool()")
	callStack.ShowStack()
	callStack.Pop()
	callStack.ShowStack()
	callStack.Pop()
	callStack.ShowStack()
	fmt.Println("Finished")
}
