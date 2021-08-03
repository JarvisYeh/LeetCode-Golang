package main

import "math"

// MIDTERM_maxPathSum
// maximum sum of path from one leaf node to another leaf node
func MIDTERM_maxPathSum(root *TreeNode) int {
	max := new(int)
	*max = math.MinInt64
	MIDTERM_helper(root, max)
	return *max
}

/** Note for one child situation
 *      root
 *     /   \
 *    nil  root.Right
 */
func MIDTERM_helper(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := MIDTERM_helper(root.Left, max)
	right := MIDTERM_helper(root.Right, max)

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
