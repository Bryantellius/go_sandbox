package algos

import (
	"fmt"
)

// Write a program to implement a buffer (FIFO)

// requirements
// 1. it represents a _queue_, or First In, First Out data structure
// 2. items added to the queue are appended to the end, and items removed are taken from the start

type Node struct {
	data int
	next *Node
}

func CreateNode(data int) (n *Node) {
	newNode := Node{data: data}
	return &newNode
}

type Queue struct {
	front, rear *Node
}

// Adds an element to the end of the queue
func (q *Queue) Enqueue(value int) {
	fmt.Printf("Enqueueing %d...\n", value)
	newNode := CreateNode(value)
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	}

	
	q.rear.next = newNode
	q.rear = newNode
}

// Removes an element from the front of the queue
func (q *Queue) Dequeue() (value int) {
	if q.front == nil {
		return -1
	}

	value = q.front.data
	fmt.Printf("Dequeueing %d...", value)
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	return
}

// Returns the element at the beginning of the start queue without removing it
func (q *Queue) Peak() (value int) {
	value = q.front.data
	return
}

// Returns the element at the beginning of the end queue without removing it
func (q *Queue) Rear() (value int) {
	value = q.rear.data
	return
}

func CreateQueue() (q *Queue) {
	return &Queue{}
}
