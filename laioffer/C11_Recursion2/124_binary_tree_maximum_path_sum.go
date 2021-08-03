package main

import "math"

func maxPathSum124(root *TreeNode) int {
	max := new(int)
	*max = math.MinInt64
	helper124(root, max)
	return *max
}

// helper returns the max path sum start from root downward
func helper124(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := helper124(root.Left, max)
	right := helper124(root.Right, max)

	// curr is the max path sum in current root subtree
	curr := root.Val
	if left > 0 {
		curr += left
	}
	if right > 0 {
		curr += right
	}
	if curr > *max {
		*max = curr
	}

	// returns path sum max starts from root and downward
	// in current root subtree
	if left > right && left > 0 {
		return root.Val + left
	}
	if right > left && right > 0 {
		return root.Val + right
	}
	return root.Val
}
