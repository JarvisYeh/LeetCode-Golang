package main

import "math"

func maxDiff(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := maxDiff(root.Left, max)
	right := maxDiff(root.Right, max)
	if int(math.Abs(float64(left)-float64(right))) > *max {
		*max = int(math.Abs(float64(left) - float64(right)))
	}

	// return total number of nodes at root
	return left + right + 1
}
