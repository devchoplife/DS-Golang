package main

import "fmt"

var countNodes int

type Node struct {
	Key   int
	left  *Node
	right *Node
}

// Insert will add a node to the tree
// the key to add should not exist in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.right == nil {
			n.right = &Node{Key: k}
		} else {
			n.right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.left == nil {
			n.left = &Node{Key: k}
		} else {
			n.left.Insert(k)
		}
	}
}

// Search will take in akey value and return true if there is a node with that value
func (n *Node) Search(k int) bool {

	countNodes++
	// if there is no match
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.right.Search(k)
	} else if n.Key > k {
		//move left
		return n.left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)
	tree.Insert(135)
	tree.Insert(288)
	tree.Insert(28)

	//fmt.Println(tree.right)
	//fmt.Println(tree.right.right)
	fmt.Println(tree.Search(18))
	fmt.Println(countNodes)

}
