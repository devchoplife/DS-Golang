package main

/*
One of the most basic data structures and it is commonly asked in interviews

e.g. Design a hashmap without using any builtin library

Data is stored in a has table ins a way that the data is passed through a hash function and a has code is generated
The data is stored along with the hash code so when you want to search for the string, it simply looks for the has code
and returns the strin.

Hash Function
==========================
The hash function gets the ASCII code of the characters in the word and addsthem together
Then divides the sum by 100 and use the remainder as the hash code
This way the has code is always between 0 and 100

Collision handling
============================
There is a chance that two names might have exactly the same hash code. Handling this type of situation is known as
Collision handling:

There are two ways of resolving collision handling namely:
++ Open adressing
++separate chaining

Open addressing
+++++++++++++++++++++++++
Open addressng stores the data in the next index if the original hash code generated is already taken
But with open addressing, as the array gets filled up the data gets further and further away from where it should be
and we will start to lose the benefits of hash tables

Separate chaining
============================
basically, separate chaining storing multiple names in one index using linked-lists.
A combination of array and linked lists makes separate chaining a very fast and effecient structure for has tables


Hash tables are used for storing key and value pairs
The names are caled keys
Linked lists in the array index are called Buckets

The time complexity of has tables for inserting, deleting and searchig is expected to be constant O(1)
The worst case would be when all keys to be stored collide then the table becomes more like a linked list which makes
insertion, deletion and searching toime consuming  but this is a rare scenario
*/
