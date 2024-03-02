package algos

import "fmt"

// Write a program to manage an array of many stacks.

type StackHandler struct {
	stacks []*Stack
}

// Creates a new Stack Handler.
func CreateStackHandler() *StackHandler {
	return &StackHandler{}
}

// Iterates through the stacks in the Stack Handler and process each stack and their respective stack nodes.
func (sh *StackHandler) ProcessStacks() {
	fmt.Println("Starting all stack processes")
	var counter int
	for len(sh.stacks) > 0 {
		sh.processStack(sh.Shift())
		counter++
	}
	fmt.Printf("Finished processing %d stack(s)", counter)
}

// Adds a new stack to the end of the Stack handler.
func (sh *StackHandler) Push(s *Stack) {
	sh.stacks = append(sh.stacks, s)
}

// Removes the first stack in the stack handler.
func (sh *StackHandler) Shift() *Stack {
	removed := sh.stacks[0]
	sh.stacks = sh.stacks[1:]
	return removed
}

// Iterates through each stack node of a stack.
func (sh *StackHandler) processStack(s *Stack) {
	fmt.Println("Starting new stack process")
	var node string
	for {
		s.ShowStack()
		node = s.Pop()
		if node == "" {
			break
		}
	}
	fmt.Println("Finished processing stack")
}
