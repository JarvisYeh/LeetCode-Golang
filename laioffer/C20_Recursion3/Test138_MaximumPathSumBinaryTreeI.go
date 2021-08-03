package main

import "math"

// maximum sum of path from one leaf node to another leaf node
// same as mid-term question
func maxPathSum138(root *TreeNode) int {
	max := new(int)
	*max = math.MinInt64
	helper138(root, max)
	return *max
}

/** Note for one child situation
 *      root
 *     /   \
 *    nil  root.Right
 */
func helper138(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := helper138(root.Left, max)
	right := helper138(root.Right, max)

	// only update max when root has two children
	// otherwise root + left + right is not a valid leaf to leaf sum
	if root.Left != nil && root.Right != nil && root.Val+left+right > *max {
		*max = root.Val + left + right
	}

	// if root has only one child, return one side + root.Val
	if root.Left == nil {
		return root.Val + right
	}

	if root.Right == nil {
		return root.Val + left
	}

	// if has two children, return the max size + root.Val
	if left > right {
		return root.Val + left
	} else {
		return root.Val + right
	}
}
