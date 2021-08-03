package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{nil, nil, 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index+1 < 0 || index+1 > this.size {
		return -1
	}

	head := this.head
	for i := 0; i < index; i++ {
		head = head.next
	}
	return head.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{val, nil}
	if this.size == 0 {
		this.head = newNode
		this.tail = newNode
	} else {
		newNode.next = this.head
		this.head = newNode
	}
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{val, nil}
	if this.size == 0 {
		this.head = newNode
		this.tail = newNode
	} else {
		this.tail.next = newNode
		this.tail = newNode
	}
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	// corner cases
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}

	prev := this.head
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}
	newNode := &Node{val, prev.next}
	prev.next = newNode
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	// corner cases
	if index < 0 || index > this.size-1 {
		return
	}

	// delete only one element
	if this.size == 1 {
		this.head = nil
		this.tail = nil
		this.size--
		return
	}

	// delete first element, head will change
	if index == 0 {
		this.head = this.head.next
		this.size--
		return
	}

	// delete last element, tail will change
	if index == this.size-1 {
		prev := this.head
		for prev.next != this.tail {
			prev = prev.next
		}
		prev.next = nil
		this.tail = prev
		this.size--
		return
	}

	// delete element in the middle
	prev := this.head
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
	this.size--
}

func (this *MyLinkedList) printMyLinkedList() {
	curr := this.head
	for curr != nil {
		fmt.Print(strconv.Itoa(curr.val) + "->")
		curr = curr.next
	}
	fmt.Println("")
}

func main() {
	list := Constructor()
	list.AddAtHead(9)
	list.printMyLinkedList()

	fmt.Println(list.Get(1))
	list.printMyLinkedList()

	list.AddAtIndex(1, 1)
	list.printMyLinkedList()

	list.AddAtIndex(1, 7)
	list.printMyLinkedList()

	list.DeleteAtIndex(1)
	list.printMyLinkedList()

	list.AddAtHead(7)
	list.printMyLinkedList()

	list.AddAtHead(4)
	list.printMyLinkedList()

	list.DeleteAtIndex(1)
	list.printMyLinkedList()

	list.AddAtIndex(1, 4)
	list.printMyLinkedList()

	list.AddAtHead(2)
	list.printMyLinkedList()

	list.DeleteAtIndex(5)
	list.printMyLinkedList()
}
