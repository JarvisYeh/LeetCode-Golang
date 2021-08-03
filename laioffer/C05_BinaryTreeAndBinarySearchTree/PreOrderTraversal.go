package main

func preorderTraversal144(root *TreeNode) []int {
	var res []int
	preOrderRecursion144(root, &res)
	return res
}

func preOrderRecursion144(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrderRecursion144(root.Left, res)
	preOrderRecursion144(root.Right, res)
}

func preOrderIter144(root *TreeNode, res *[]int) {
	stack := make([]*TreeNode, 0)
	if root == nil {
		return
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		*res = append(*res, curr.Val)
		stack = stack[:len(stack)-1]
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
}
