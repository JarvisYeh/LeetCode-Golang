package main

import "fmt"

type Node146 struct {
	Key   int
	Value int
	Prev  *Node146
	Next  *Node146
}

type LRUCache146 struct {
	Head     *Node146
	Tail     *Node146
	Capacity int
	Size     int
	NodeMap  map[int]*Node146
}

func Constructor(capacity int) LRUCache146 {
	lru := LRUCache146{nil, nil, capacity, 0, map[int]*Node146{}}
	return lru
}

func (this *LRUCache146) Get(key int) int {
	if node, ok := this.NodeMap[key]; !ok {
		return -1
	} else {
		// update position of node in list
		// get won't involve with capacity limitation
		this.remove(node)
		this.add(node)
		return node.Value
	}
}

func (this *LRUCache146) Put(key int, value int) {
	if node, ok := this.NodeMap[key]; ok { // if LRU contains key
		node.Value = value
		this.Get(key)
	} else if this.Size == this.Capacity { // if LRU doesn't contains key and capacity exceed limit
		this.remove(this.Head)
		this.Put(key, value)
	} else {
		newNode := &Node146{key, value, nil, nil}
		this.add(newNode)
	}
	fmt.Println("NULL")
}

// remove a node from lru
// 1. remove from hashMap
// 2. remove from linked list
func (this *LRUCache146) remove(node *Node146) {
	// delete from hashMap
	delete(this.NodeMap, node.Key)
	// remove from linked list
	prev, next := node.Prev, node.Next
	if prev == nil && next == nil { // remove the only element
		this.Head = nil
		this.Tail = nil
	} else if prev == nil { // node == head
		next.Prev = nil
		this.Head = next
	} else if next == nil { // node == tail
		prev.Next = nil
		this.Tail = prev
	} else { // some middle node
		prev.Next = next
		next.Prev = prev
	}
	// safe deletion
	node.Prev = nil
	node.Next = nil
	this.Size--
}

// add a node to LRU
// 1. append to hashMap
// 2. append to linked list tail
func (this *LRUCache146) add(node *Node146) {
	this.NodeMap[node.Key] = node
	if this.Size == 0 {
		this.Head = node
		this.Tail = node
	} else {
		node.Prev = this.Tail
		this.Tail.Next = node
		this.Tail = node
	}
	this.Size++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lru := Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	for i := 4; i > 0; i-- {
		fmt.Printf("GET(%d): %d]\n", i, lru.Get(i))
	}
	lru.Put(5, 5)
	for i := 1; i <= 5; i++ {
		fmt.Printf("GET(%d): %d]\n", i, lru.Get(i))
	}
}
