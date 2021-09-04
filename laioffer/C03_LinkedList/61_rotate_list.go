package main

func rotateRight61(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	tail, curr := head, head
	n := 0
	for curr != nil {
		n++
		tail = curr
		curr = curr.Next
	}

	// rotate right k steps, same as rotate right k % n steps
	// so k % n nodes from end will be rotated to front
	// the left unmoved nodes are (n - k % n)
	leftSize := n - (k % n)
	if leftSize == 0 || leftSize == n {
		return head
	}
	var prev *ListNode
	curr = head
	for i := 0; i < leftSize; i++ {
		prev = curr
		curr = curr.Next
	}

	tail.Next = head
	prev.Next = nil
	return curr
}
