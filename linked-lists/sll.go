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

// print all the data in the list
func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	// handling special cases so we do not get segmentation violation error/runtime error

	// empty list
	if l.length == 0 {
		return
	}

	// value to delete is the head
	if value == l.head.data {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head

	// if value does not exist
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
	}

	for previousToDelete.next.data != value {
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}

	// defined nodes prepend to list
	node1 := &node{data: 23}
	node2 := &node{data: 24}
	node3 := &node{data: 25}
	node4 := &node{data: 26}
	node5 := &node{data: 27}
	node6 := &node{data: 28}

	// prepend nodes
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	// before deleting
	myList.printListData()

	myList.deleteWithValue(27)
	// After deleting
	myList.printListData()
}
