package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

// LinkedList : the list object
type LinkedList struct {
	head *node
	tail *node
	size int
}

// Create : Creates a linked list
func Create() *LinkedList {
	newList := LinkedList{head: nil, size: 0, tail: nil}
	return &newList
}

// InsertAfter : inserts data into list
func InsertAfter(list *LinkedList, data int) *LinkedList {
	newNode := createNode(data)

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}

	list.size++
	return list
}

// InsertBefore : inserts new node to front of list
func InsertBefore(list *LinkedList, data int) *LinkedList {
	newNode := createNode(data)

	if list.head == nil {
		list.head = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}

	list.size++
	return list
}

// PrintList : prints list in human readable format
func PrintList(list *LinkedList) {
	printNode(list.head)
	fmt.Println()
}

func printNode(node *node) {
	fmt.Printf("%d ", node.data)

	if node.next != nil {
		printNode(node.next)
	}
}

func createNode(data int) *node {
	newNode := node{data: data, next: nil}
	return &newNode
}
