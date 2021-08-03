package main

import "fmt"

// iteration
func reverseKGroup25(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{-1, head}
	// prev -> gH ... gT -> next
	var prev, next, gH, gT *ListNode
	prev = dummyHead

	for true {
		gH = prev.Next
		gT = prev
		for i := 0; i < k; i++ {
			if gT.Next != nil {
				gT = gT.Next
			} else {
				return dummyHead.Next
			}
		}
		// now gH is the head node of the group
		// gT is the tail node of the group

		// record the next node after the group
		next = gT.Next
		// reverse the group, after reverse, newHead = gT, newTail = gH
		reverse25(gH, gT)
		// connect prev, next to the group newHead and newTail
		prev.Next = gT
		gH.Next = next

		// update prev pointer
		prev = gH
	}
	return nil
}

func reverse25(head, tail *ListNode) {
	endNode := tail.Next
	curr := head
	var prev *ListNode
	for curr != endNode {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}

/**
 * Recursion method
 * getNewHead of next group
 * reverse current group, append nextNewHead to tail
 * return newHead of current group
 */
func reverseKGroup25II(head *ListNode, k int) *ListNode {
	// base case: left node is less than k node
	curr := head
	for i := 0; i < k; i++ {
		if curr == nil {
			return head
		}
		curr = curr.Next
	}
	// now curr is the start node of next group

	// complete the reversing for following nodes, and retrive the new head of next group
	nextHead := reverseKGroup25II(curr, k)

	// reverse current group
	prev, curr := nextHead, head
	for i := 0; i < k; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// now return the new head of current group
	return prev
}

func main() {
	head := CreateListWithArray([]int{1, 2, 3, 4, 5})
	fmt.Println(reverseKGroup25(head, 2))
}
