package main

import "math"

// max sum of sub-path from root to leaf

// method 1:
// pass value from top to bottom
// with DP (largest subarray sum)
// TC: O(n)
// SC: O(h)
func maxPathSum140I(root *TreeNode) int {
	max := new(int)
	*max = math.MinInt64
	helper140I(root, 0, max)
	return *max
}

func helper140I(root *TreeNode, prefixSum int, max *int) {
	if root == nil {
		return
	}

	// update prefixSum with curr level root
	if prefixSum < 0 {
		prefixSum = root.Val
	} else {
		prefixSum += root.Val
	}

	// update max
	if prefixSum > *max {
		*max = prefixSum
	}

	helper140I(root.Left, prefixSum, max)
	helper140I(root.Right, prefixSum, max)
}

// method 2:
// pass value from bottom to top
// return to parent node: max sum of sub-path start at root
// TC: O(n)
// SC: O(h)
func maxPathSum140II(root *TreeNode) int {
	max := new(int)
	*max = math.MinInt64
	helper140II(root, max)
	return *max
}

func helper140II(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := helper140II(root.Left, max)
	right := helper140II(root.Right, max)

	// curr = max(max(left, right), 0)
	curr := root.Val
	if left > 0 && right > 0 {
		if left > right {
			curr += left
		} else {
			curr += right
		}
	} else if left > 0 {
		curr += left
	} else if right > 0 {
		curr += right
	}

	// update max
	if curr > *max {
		*max = curr
	}
	return curr
}
