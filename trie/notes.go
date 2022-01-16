package main

/*
A trie is a tree structure which is mostly used for storing words
Each node consist of a word

IT helps us search words faster
One great applicationn of tries is auto-complete

Tries have a root node at the top which have children nodes, one for each possible alphabet
Each node needs to hold an array containing the addresses of the next node
Also, the node needs to have a boolean value indicating if the node is the end of the word or not

Time Complexity is O(m) where m is the length of the word being searched

A trie can take a lot of space because the trie gets larger as we ad more words
This is why tries are known to have a trade-off between time and space
*/
