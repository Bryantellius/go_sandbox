package main

import (
	"fmt"

	"github.com/Bryantellius/go_sandbox/algos"
	getgrid "github.com/Bryantellius/go_sandbox/get_grid"
)

func main() {
	file1 := "./data/cera_getGrid_output_without_change.csv"
	file2 := "./data/cera_getGrid_output_with_change.csv"
	outputFile := "./data/cera_getGrid_out_distinct.csv"
	getgrid.CompareCSVsAndRemoveDuplicates(file1, file2, outputFile)
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

func DoStackProcessor() {
	processor := algos.CreateStackHandler()
	stackA := algos.CreateStack()
	stackB := algos.CreateStack()
	stackC := algos.CreateStack()

	stackA.Push("A: func Apple()")
	stackA.Push("A: func Awesome()")
	stackA.Push("A: func Antagonism()")
	stackB.Push("B: func Banana()")
	stackB.Push("B: func Boring()")
	stackB.Push("B: func Belittling()")
	stackB.Push("B: func Brute()")
	stackB.Push("B: func Benefit()")
	stackC.Push("C: func Carrot()")
	stackC.Push("C: func Capture()")

	processor.Push(stackA)
	processor.Push(stackB)
	processor.Push(stackC)

	processor.ProcessStacks()
}

func DoQueue() {
	fmt.Println("Starting")
	queue := algos.CreateQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Peak())                         // should be 1
	fmt.Println(queue.Rear())                         // should be 3
	fmt.Printf("Processing %d...\n", queue.Dequeue()) // should be 1
	fmt.Println(queue.Peak())                         // should be 2
	queue.Enqueue(4)
	fmt.Println(queue.Rear()) // should be 4
	fmt.Println("Finished")
}

func DoBatch() {
	algos.RunBatch("./data/commands.txt")
}
