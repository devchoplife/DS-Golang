package main

import "fmt"

// struct to hold the queue
type Queue struct {
	items []int
}

// Enqueue adds items to the end of a queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue removes a value at the begining of the queue
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)

	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
