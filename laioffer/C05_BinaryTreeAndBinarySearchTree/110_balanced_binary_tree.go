package main

import "math"

func isBalanced110(root *TreeNode) bool {
	if isBalancedRecursion110(root) == -1 {
		return false
	} else {
		return true
	}
}

func isBalancedRecursion110(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := isBalancedRecursion110(root.Left)
	right := isBalancedRecursion110(root.Right)

	// if subtree already unbalanced, return -1
	if left == -1 || right == -1 {
		return -1
	}

	// if root as tree is unbalanced, return -1
	if math.Abs(float64(right-left)) > 1 {
		return -1
	}

	// else return height of root as tree
	if right > left {
		return right + 1
	} else {
		return left + 1
	}
}
