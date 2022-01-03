package main

import "fmt"

// struct holding stack slice
type Stack struct {
	items []int
}

// Push -> add/append data to the end of a stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// op -> remove data from the end of a stack
// also returns the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	myStack := Stack{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(345)

	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
