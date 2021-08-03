package main

import "fmt"

func insert708(aNode *ListNode, x int) *ListNode {
	// corner case
	if aNode == nil {
		newHead := &ListNode{x, nil}
		newHead.Next = newHead
		return newHead
	}

	// one pass to find the maximal
	// minimal is maximal.Next
	maxNode := aNode
	for maxNode.Next != aNode {
		if maxNode.Val > maxNode.Next.Val { // since sorted list
			break
		}
		maxNode = maxNode.Next
	}
	minNode := maxNode.Next

	// if x < min or x > max, insert as ... max -> x -> min ...
	if x < minNode.Val || x > maxNode.Val {
		newNode := &ListNode{x, minNode}
		maxNode.Next = newNode
		return aNode
	}

	// if  min < x < max
	// start from minNode, find the place to insert the newNode
	// stop at: curr -> x -> curr.Next
	curr := minNode
	for curr.Next != minNode && curr.Next.Val < x {
		curr = curr.Next
	}
	// insert that node after curr
	newNode := &ListNode{x, curr.Next}
	curr.Next = newNode
	return aNode
}

func main() {
	head := CreateListWithArray([]int{3, 4, 1})
	curr := head
	for curr = head; curr.Next != nil; curr = curr.Next {
	}
	curr.Next = head
	newHead := insert708(head, 2)
	curr = newHead
	for curr.Next != newHead {
		fmt.Printf("%v -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println(curr.Val)
}
