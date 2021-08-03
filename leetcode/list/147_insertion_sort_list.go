package main

// TC: O(n^2)
// SC: O(1)
func insertionSortList(head *ListNode) *ListNode {
	// newHead is the head of new created list, start from size = 0
	// every time insert the node from original list into that new list
	// still in place operation
	newDummyHead := &ListNode{-1, nil}

	// prev is the position in the new linked list, position before the node to be inserted
	// curr is the node in the original list to be checked position
	curr := head
	// every time find the place to insert curr node in new list
	for curr != nil {
		prev := newDummyHead
		for prev.Next != nil && prev.Next.Val < curr.Val {
			prev = prev.Next
		}

		next := curr.Next
		// insert the node after prev node in new list
		curr.Next = prev.Next
		prev.Next = curr
		// move forward the pointer in the original list
		curr = next
	}
	return newDummyHead.Next
}
