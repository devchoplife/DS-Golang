package main

import "fmt"

// ArraySize is the size of the hash table array
const ArraySize = 7

// Hashtable will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the hash table
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list node that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

// Init wil create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// HASH TABLE
// ========================
// Insert will take in a new key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and return true if the key is stored in the hash table

// Delete will take in a key and delete it from the hash table

// BUCKET
// ========================
// insert
func (b *bucket) insert(k string) {
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

// search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func main() {
	testHashTable := Init()
	fmt.Println(testHashTable)

	testBucket := &bucket{}
	testBucket.insert("RANDY")
	fmt.Println(testBucket.search("RANDY"))
	fmt.Println(testBucket.search("ERIC"))
}
