package main

import "fmt"

// Queue represents a queue data structure
type Queue struct {
	data []interface{}
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.data = append(q.data, item)
}

// Dequeue removes and returns the first element from the queue
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item
}

// Front returns the first element in the queue without removing it
func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	return q.data[0]
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.data)
}

// Example usage
func main() {
	queue := Queue{}

	// Enqueue elements to the queue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Dequeue and print elements from the queue
	for !queue.IsEmpty() {
		item := queue.Dequeue()
		fmt.Println(item)
	}
}
