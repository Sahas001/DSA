package main

import (
	"fmt"
)

type Node struct {
	prev  *Node
	value int
	next  *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) InsertAtHead(m int) {
	node := &Node{nil, m, nil}
	// Check if the linked list is empty at first
	if linkedList.headNode != nil {
		linkedList.headNode.prev = node
		node.next = linkedList.headNode
	}
	linkedList.headNode = node
}

func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.next {
		fmt.Printf("%d\t", node.value)
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.next {
		lastNode = node
	}
	return lastNode
}

func (linkedList *LinkedList) InsertAtEnd(m int) {
	lastNode := linkedList.LastNode()
	node := &Node{nil, m, nil}
	node.prev = lastNode
	lastNode.next = node
}

func (linkedList *LinkedList) ReturnNode(m int) *Node {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.next {
		if node.value == m {
			return node
		}
	}
	return node
}

func (linkedList *LinkedList) InsertAfter(nodeM int, m int) {
	ourNode := &Node{nil, m, nil}
	requiredNode := linkedList.ReturnNode(nodeM)
	if requiredNode.next != nil {
		ourNode.next = requiredNode.next
		ourNode.prev = requiredNode
		requiredNode.next = ourNode
	}
}

func main() {
	linkedList := &LinkedList{}
	linkedList.InsertAtHead(1)
	linkedList.InsertAtHead(5)
	linkedList.InsertAtHead(3)
	linkedList.InsertAtEnd(7)
	linkedList.IterateList()
	fmt.Printf("\n")
	linkedList.InsertAfter(3, 8)
	linkedList.IterateList()
}
