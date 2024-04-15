package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) InsertAtHead(m int) {
	node := Node{}
	node.value = m
	// if node.next != nil {
	node.next = linkedList.headNode
	// }
	linkedList.headNode = &node
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
		if node.next == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (linkedList *LinkedList) InsertAtEnd(m int) {
	node := &Node{m, nil}
	var lastNode *Node
	lastNode = linkedList.LastNode()
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
	ourNode := &Node{m, nil}
	requiredNode := linkedList.ReturnNode(nodeM)
	if requiredNode.next != nil {
		ourNode.next = requiredNode.next
		requiredNode.next = ourNode
	}
}

func main() {
	linkedList := &LinkedList{}
	linkedList.InsertAtHead(1)
	linkedList.InsertAtHead(2)
	linkedList.InsertAtHead(4)
	linkedList.InsertAtEnd(5)
	linkedList.IterateList()
	linkedList.InsertAfter(2, 3)
	fmt.Printf("\n")
	linkedList.IterateList()
}
