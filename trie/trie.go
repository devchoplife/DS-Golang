package main

import "fmt"

// AlphabetSize is the number of possibe characters in the node
const AlphabetSize = 26

// Node struct - Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool // to indicate if it is end of the word or not
}

// Trie struct
type Trie struct {
	root *Node // basically a pointer pointing to the root node
}

// initTrie will create a new trie
func initTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wLen := len(w)
	currentNode := t.root
	for i := 0; i < wLen; i++ {
		// charIndex gets the index of the character using the ascii index
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

// Search will take in a word and return true if that word is included in the trie
func (t *Trie) Search(w string) bool {
	wLen := len(w)
	currentNode := t.root
	for i := 0; i < wLen; i++ {
		// charIndex gets the index of the character using the ascii index
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main() {
	myTrie := initTrie()
	toAdd := []string{
		"mine",
		"yours",
		"ours",
		"theirs",
		"they",
		"he",
		"her",
		"here",
	}

	for _, w := range toAdd {
		myTrie.Insert(w)
	}
	fmt.Println(myTrie.Search("theirs"))
	fmt.Println(myTrie.Search("potato"))
}
