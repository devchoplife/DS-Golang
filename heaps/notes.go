package heaps

/*
Heaps are especially good for implementing priority queues

In priority queues, the one with the highest priority is taken out first

A heap is a data structure that can be expressed as a complete tree
A complete tree means that all the levels in the tree are full except for the lowest
level where some nodes might be empty but has all its nodes to the left.

BINARY HEAP
========================
The nodes have up to two children here

MAX HEAP
The larget key will be stored in the root and for all the nodes in the tree, every parent will
have a higher key than its children
It would be useful when we need to repeatedly remove a key that is the largest because the
largest is always at the root

MIN HEAP
=========================
Here the root is the smallest key

Whenever we insert into a heap, we insert into the bottom right node.
When a new node is added is compared to its key to see if it is larger (Max heap), if larger it switches position with its key until it gets to a node where
it is smaller than theke. THis process is called HEAPIFY

Heapify is also needed after taking an item off the heap

WHen the root is removed, the last node is inserted in the root and then the swapping process begins from the top

Time complexity for inserting and extracting in a heap is O(logn)

The left child is always an odd number and the right child is always an even numberand they are both going to indicate
te same parent index
*/
