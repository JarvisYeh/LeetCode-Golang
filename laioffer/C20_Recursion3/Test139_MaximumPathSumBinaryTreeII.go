package main

import "math"

// maximum sum of path from any node to any node
func maxPathSum139(root *TreeNode) int {
	max := new(int)
	*max = math.MinInt64
	helper139(root, max)
	return *max
}

// helper return max sum of path from root downwards (including root)
// current level update max with any node to any node sum of path (including root node)
func helper139(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := helper139(root.Left, max)
	right := helper139(root.Right, max)

	curr := root.Val
	if left > 0 {
		curr += left
	}
	if right > 0 {
		curr += right
	}
	// update max
	if curr > *max {
		*max = curr
	}

	// return to previous level
	if left < 0 && right < 0 { // l < 0, r < 0
		return root.Val
	} else if left < 0 { // l < 0, r > 0
		return root.Val + right
	} else if right < 0 { // l > 0, r < 0
		return root.Val + left
	}
	// now left > 0 && right > 0
	if left > right {
		return left + root.Val
	} else {
		return right + root.Val
	}
}
