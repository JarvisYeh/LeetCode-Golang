package main

type TreeNodeLeft struct {
	Val          int
	NumNodesLeft int
	Left         *TreeNodeLeft
	Right        *TreeNodeLeft
}

func setNumNodesLeft646(root *TreeNodeLeft) int {
	if root == nil {
		return 0
	}

	left := setNumNodesLeft646(root.Left)
	right := setNumNodesLeft646(root.Right)
	root.NumNodesLeft = left
	return left + right + 1
}
