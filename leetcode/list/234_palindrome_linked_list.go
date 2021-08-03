package main

func isPalindrome(head *ListNode) bool {
	mid := findMiddle(head)
	two := mid.Next
	mid.Next = nil
	two = reverseList(two)
	for two != nil {
		if head.Val != two.Val {
			return false
		}
		two = two.Next
		head = head.Next
	}
	return true
}
