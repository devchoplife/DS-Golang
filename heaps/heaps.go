package main

import "fmt"

// MaxHeap struct
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	// Storing the largest key and returning it first
	extracted := h.array[0]

	// prevent error that might arise when extracting from empty array
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	// make the last index the root index
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]

	// remove the last index
	h.array = h.array[:lastIndex]
	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp will heapif from bpttom to the top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDwon will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	// loop while index has at least one child
	for l <= lastIndex {
		// when left child is the only child
		// ensure you always put this first because if there is only one child we do not want to be accessing the right
		// child
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Get the parent index
// Parent index can be expreesed as index minus 1 divided by 2
func parent(i int) int {
	return (i - 1) / 2
}

// Get the left child index
func left(i int) int {
	return (2 * i) + 1
}

// Get the right child index
func right(i int) int {
	return (2 * i) + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17, 19}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

}
