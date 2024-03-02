package algos

import (
	"fmt"
)

// Write a program to create a stack (LIFO) which can store any number of local variables for each subroutine.
// Hint: use a linked list for the stack and local variables

type StackNode struct {
	value string
	next  *StackNode
}

type Stack struct {
	stackPointer *StackNode
}

func CreateStack() (s *Stack) {
	s = &Stack{stackPointer: &StackNode{}}
	return
}

func (s *Stack) Push(signature string) {
	fmt.Printf("Pushing %s...\n", signature)
	next := s.stackPointer
	s.stackPointer = &StackNode{value: signature, next: next}
}

func (s *Stack) Pop() string {
	val := s.stackPointer.value
	fmt.Printf("Popping %s...\n", val)
	if s.stackPointer.next != nil {
		s.stackPointer = s.stackPointer.next
	}

	return val
}

func (s *Stack) ShowStack() {
	p := s.stackPointer
	level := 0

	for {
		fmt.Printf("| %d: %s |\n", level, p.value)
		p = p.next
		level++
		if p == nil {
			break
		}
	}
}
