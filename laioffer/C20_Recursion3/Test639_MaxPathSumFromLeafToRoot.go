package main

import "math"

// max sum of path from root to leaf

// method1 : pass {sum of root to curr.parent} downward
func maxPathSum639I(root *TreeNode) int {
	max := new(int)
	*max = math.MinInt64
	helper639(root, 0, max)
	return *max
}

func helper639(root *TreeNode, currSum int, max *int) {
	if root == nil {
		return
	}

	// update max only in leaf node but not in nil node
	if root.Left == nil && root.Right == nil {
		if currSum+root.Val > *max {
			*max = currSum + root.Val
		}
	}
	helper639(root.Left, currSum+root.Val, max)
	helper639(root.Right, currSum+root.Val, max)
}

// method2 : pass {max sum from curr to leaf} upward
func maxPathSum639II(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxPathSum639II(root.Left)
	right := maxPathSum639II(root.Right)
	// for single children node
	if root.Left == nil {
		return right + root.Val
	}
	if root.Right == nil {
		return left + root.Val
	}
	// for two children node
	if left > right {
		return left + root.Val
	} else {
		return right + root.Val
	}
}
