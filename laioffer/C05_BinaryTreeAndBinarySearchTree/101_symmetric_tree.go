package main

func isSymmetric(root *TreeNode) bool {
	return isSymmetricRecursion(root.Left, root.Right)
}

func isSymmetricRecursion(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricRecursion(left.Left, right.Right) && isSymmetricRecursion(left.Right, right.Left)
}
