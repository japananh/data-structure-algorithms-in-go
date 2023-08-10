package main

import (
	"errors"
	"fmt"
)

var ErrQueueIsFull = errors.New("queue is full")
var ErrQueueIsEmpty = errors.New("queue is empty")

var defaultCapacity = 1000000

// queue represents a queue data structure
type queue struct {
	front    int
	rear     int
	size     int
	capacity int
	data     []interface{}
}

type Queue interface {
	Enqueue(item interface{}) (interface{}, error)
	Dequeue() (interface{}, error)
	Front() (interface{}, error)
	Rear() (interface{}, error)
	Size() int
	Capacity() int
	IsEmpty() bool
	IsFull() bool
}

// NewQueue creates a new queue
func NewQueue(capacity int) Queue {
	if capacity < defaultCapacity {
		capacity = defaultCapacity
	}

	return &queue{
		capacity: capacity,
		rear:     capacity - 1,
		data:     make([]interface{}, capacity),
	}
}

// Enqueue adds an element to the end of the queue
func (q *queue) Enqueue(item interface{}) (interface{}, error) {
	if q.IsFull() {
		return nil, ErrQueueIsFull
	}

	q.size++
	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = item

	return item, nil
}

// Dequeue removes and returns the first element from the queue
func (q *queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueIsEmpty
	}

	item := q.data[q.front]
	q.size--
	q.data[q.front] = nil
	q.front = (q.front + 1) % q.capacity

	return item, nil
}

// Front returns the first element in the queue without removing it
func (q *queue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueIsEmpty
	}
	return q.data[q.front], nil
}

// Rear returns the last element in the queue without removing it
func (q *queue) Rear() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueIsEmpty
	}
	return q.data[q.rear], nil
}

// Size returns the number of elements in the queue
func (q *queue) Size() int {
	return q.size
}

// Capacity returns the capacity of the queue
func (q *queue) Capacity() int {
	return q.capacity
}

// IsEmpty checks if the queue is empty
func (q *queue) IsEmpty() bool {
	return q.size == 0
}

// IsFull checks if the queue is full
func (q *queue) IsFull() bool {
	return q.size == q.capacity
}

// Example usage
func main() {
	// Create a queue
	queue := NewQueue(3)

	// Enqueue elements to the queue
	if _, err := queue.Enqueue(10); err != nil {
		fmt.Printf("Error when enqueueing: %v\n", err)
	}
	if _, err := queue.Enqueue("20 is a number"); err != nil {
		fmt.Printf("Error when enqueueing: %v\n", err)
	}
	if _, err := queue.Enqueue(30.1); err != nil {
		fmt.Printf("Error when enqueueing: %v\n", err)
	}
	if _, err := queue.Enqueue(40); err != nil {
		fmt.Printf("Error when enqueueing: %v\n", err)
	}

	// Dequeue and print elements from the queue
	for !queue.IsEmpty() {
		item, err := queue.Dequeue()
		if err != nil {
			fmt.Printf("Error when dequeueing %v\n", err)
		}
		fmt.Printf("Deque item `%v` sucessfully\n", item)
	}
}
