package main

import "math"

func isValidBST98(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBST98I(root, math.MinInt64, math.MaxInt64)
}

// method 1: use range (min, max) to check a node
func isBST98I(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return isBST98I(root.Left, min, root.Val) && isBST98I(root.Right, root.Val, max)
}

// method 2: inorder traversal to see if the value increases monotonically
func isBSTII(root *TreeNode, lastSeen *int) bool {
	if root == nil {
		return true
	}

	// in-order traversal
	if !isBSTII(root.Left, lastSeen) {
		return false
	}
	if root.Val <= *lastSeen {
		return false
	}
	*lastSeen = root.Val

	return isBSTII(root.Right, lastSeen)
}
