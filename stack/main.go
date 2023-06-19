package main

import "fmt"

// Stack represents a stack data structure
type Stack struct {
	data []interface{}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

// Pop removes and returns the topmost element from the stack
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]
	return item
}

// Peek returns the topmost element from the stack without removing it
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	index := len(s.data) - 1
	return s.data[index]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.data)
}

// Example usage
func main() {
	stack := Stack{}

	// Push elements to the stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Pop and print elements from the stack
	for !stack.IsEmpty() {
		item := stack.Pop()
		fmt.Println(item)
	}
}
