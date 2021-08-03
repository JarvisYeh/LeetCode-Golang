package main

func postorderTraversal145(root *TreeNode) []int {
	var res []int
	postOrderRecursion145(root, &res)
	return res
}

func postOrderRecursion145(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	postOrderRecursion145(root.Left, res)
	postOrderRecursion145(root.Right, res)
	*res = append(*res, root.Val)
}

func postOrderIter145(root *TreeNode, res *[]int) {
	// prev is the previous visit node to curr, not the structure previous node to curr
	var prev, curr *TreeNode

	// initialize stack
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		// peek the top of stack
		curr = stack[len(stack)-1]

		// curr here is the now visit node
		// if prev is the parent of curr or prev is nil (e.g. curr is root)
		// that means it's the first time explore curr
		if prev == nil || prev.Left == curr || prev.Right == curr {
			// if curr has left child, directly route left
			// if curr only has right child, route right
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			} else if curr.Right != nil {
				stack = append(stack, curr.Right)
			} else {
				// if curr left and right both are nil
				stack = stack[:len(stack)-1]
				*res = append(*res, curr.Val)
			}
		} else if curr.Left == prev { // if prev is the left child of curr (e.g. route back from left sub tree)
			// which means nodes in left sub tree have all been stored in res
			if curr.Right != nil {
				// has right sub tree, route right
				stack = append(stack, curr.Right)
			} else {
				// does not have right sub tree, popped curr
				stack = stack[:len(stack)-1]
				*res = append(*res, curr.Val)
			}
		} else { // if prev is the right child of curr (e.g. route back from right sub tree)
			// popped curr
			stack = stack[:len(stack)-1]
			*res = append(*res, curr.Val)
		}
		prev = curr
	}
}
