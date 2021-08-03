package main

import "math"

// return the largest value that is smaller than target
func largestSmaller(root *TreeNode, target int) int {
	largest := math.MinInt64
	for root != nil {
		if root.Val >= target {
			root = root.Left
		} else { // root.Val < target, update largest
			if largest < root.Val {
				largest = root.Val
			}
			root = root.Right
		}
	}
	return largest
}
