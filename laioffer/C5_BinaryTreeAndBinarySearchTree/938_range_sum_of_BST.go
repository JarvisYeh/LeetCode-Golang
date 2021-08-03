package main

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	sumBST(root, low, high, &sum)
	return sum
}

func sumBST(root *TreeNode, low, high int, sum *int) {
	if root == nil {
		return
	}

	// prune operation
	if root.Val > low {
		sumBST(root.Left, low, high, sum)
	}

	if root.Val >= low && root.Val <= high {
		*sum = *sum + root.Val
	}

	// prune operation
	if root.Val < high {
		sumBST(root.Right, low, high, sum)
	}
}
