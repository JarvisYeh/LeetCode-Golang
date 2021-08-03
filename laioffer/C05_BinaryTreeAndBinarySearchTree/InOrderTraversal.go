package main

func inorderTraversal94(root *TreeNode) []int {
	var res []int
	inOrderRecursion94(root, &res)
	return res
}

func inOrderRecursion94(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrderRecursion94(root.Left, res)
	*res = append(*res, root.Val)
	inOrderRecursion94(root.Left, res)
}

func inOrderIter94(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	stack := make([]*TreeNode, 0)

	// initialize stack
	curr := root
	for curr != nil {
		// if curr is not null, keep pushing into stack and move left
		stack = append(stack, curr)
		curr = curr.Left
	}

	// keep popping one at a time
	// pushing all left
	for len(stack) > 0 {
		// pop one node
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		*res = append(*res, popped.Val)

		// push all left nodes in right subtree
		curr = popped.Right
		for curr != nil {
			// if curr is not null, keep pushing into stack and move left
			stack = append(stack, curr)
			curr = curr.Left
		}
	}
}
