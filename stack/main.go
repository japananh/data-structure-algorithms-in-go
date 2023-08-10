package main

import (
	"errors"
	"fmt"
)

var ErrStackIsEmpty = errors.New("stack is empty")

var defaultCapacity = 1000000

// stack represents a stack data structure
type stack struct {
	top      int
	capacity int
	data     []interface{}
}

type Stack interface {
	Push(item interface{}) (interface{}, error)
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
}

// NewStack creates a new stack
func NewStack(capacity int) Stack {
	if capacity < defaultCapacity {
		capacity = defaultCapacity
	}

	return &stack{
		capacity: capacity,
		top:      -1,
		data:     make([]interface{}, capacity),
	}
}

// Push adds an element to the top of the stack
func (s *stack) Push(item interface{}) (interface{}, error) {
	s.top = (s.top + 1) % s.capacity
	s.data[s.top] = item

	return item, nil
}

// Pop removes and returns the topmost element from the stack
func (s *stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrStackIsEmpty
	}

	item := s.data[s.top]
	s.data[s.top] = nil
	s.top = (s.top - 1) % s.capacity

	return item, nil
}

// Peek returns the top element from the stack without removing it
func (s *stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrStackIsEmpty
	}

	return s.data[s.top], nil
}

// IsEmpty checks if the stack is empty
func (s *stack) IsEmpty() bool {
	return s.top == -1
}

// Example usage
func main() {
	stack := NewStack(3)

	// Push elements to the stack
	if _, err := stack.Push(10); err != nil {
		fmt.Printf("Error when adding item in stack: %v\n", err)
	}
	if _, err := stack.Push("20 is a number"); err != nil {
		fmt.Printf("Error when adding item in stack: %v\n", err)
	}
	if _, err := stack.Push(30.1); err != nil {
		fmt.Printf("Error when adding item in stack: %v\n", err)
	}

	// Pop and print elements from the stack
	for !stack.IsEmpty() {
		item, err := stack.Pop()
		fmt.Printf("Removing item from stack: %v\n", item)
		if err != nil {
			fmt.Printf("Error when removing %v in stack: %v\n", item, err)
		}
	}
}
