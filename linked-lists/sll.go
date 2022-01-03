package main

import "fmt"

// individual nodes in the linked list
type node struct {
	data int
	next *node
}

// the whole list
type linkedList struct {
	head   *node
	length int
}

// this function adds a node to the begining of the list
func (l *linkedList) prepend(nodeToPrepend *node) {
	second := l.head
	l.head = nodeToPrepend
	l.head.next = second
	l.length++
}

func main() {
	myList := linkedList{}

	// defined nodes prepend to list
	node1 := &node{data: 23}
	node2 := &node{data: 24}
	node3 := &node{data: 25}
	node4 := &node{data: 26}

	// prepend nodes
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)

	fmt.Println(myList)
}
