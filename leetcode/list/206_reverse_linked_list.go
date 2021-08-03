package main

// iteration
func reverseList206I(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, curr *ListNode
	curr = head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

//
//recursion
func reverseList206II(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList206II(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
