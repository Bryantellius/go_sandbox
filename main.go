package main

import (
	"fmt"

	"github.com/Bryantellius/go_sandbox/algos"
)

func main() {
	DoQueue()
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

func DoQueue() {
	fmt.Println("Starting")
	queue := algos.CreateQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Peak()) // should be 1
	fmt.Println(queue.Rear()) // should be 3
	fmt.Printf("Processing %d...\n", queue.Dequeue()) // should be 1
	fmt.Println(queue.Peak()) // should be 2
	queue.Enqueue(4)
	fmt.Println(queue.Rear()) // should be 4
	fmt.Println("Finished")
}