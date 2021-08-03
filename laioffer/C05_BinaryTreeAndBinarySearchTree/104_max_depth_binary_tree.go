package main

func maxDepth104(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth104(root.Left)
	right := maxDepth104(root.Right)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
