package main

/*
A graph data structure is a collection of nodes that have data and are connected to other nodes

A graph has vertices(Vertex) and edges
When a graph has a lot of vertices it is called a DENSE graph
when the graph has relatively less edges, we say that the graph is SPARSE

It is an abstract data type. It can be implemented ad
++ Directed graph  -
++ Undirected graph - this means that every edge is bi-directional i.e. go both ways

You can also allow cycles or not
You cna also put weights on the edges e.g. time taken to pass a path from point A to B i.e. Google maps

The tree structure is also a type of graph but it has restrictions, dependin on the type of tree it is
e.g. Children cannot point to a parent, no cycles are allowed

There are two ways to represent graphs namely:
++ Adjacency List
++ Adjacency Matrix

Adjacency List
=======================
The graph is expressed as a list of vertices and each vertex will have a list holding the neighbouring vertex.
The list can be any form of list
Integer values of the vertices are known as keys and should be unique if they are going to represent a set of information

Adjacency Matrix
=======================
The data is represented in a way that they are stored in a matrux
There is the from and to side of the vertices to match the data based on weight or not
See images for more

In a sparse graph Adjacency list takes less space while for Adjacency matrix it requires space of V raised to the poweer of 2 where V
is the number of vertices

For edge lookup, the adjacency matrix is faster because we just need t do an edge lookup

*/
