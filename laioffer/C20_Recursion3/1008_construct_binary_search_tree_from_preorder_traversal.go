package main

import "math"

func bstFromPreorder1008(preorder []int) *TreeNode {
	curr := new(int)
	*curr = 0
	return build1008(preorder, curr, math.MaxInt64)
}

func build1008(preorder []int, curr *int, max int) *TreeNode {
	if *curr >= len(preorder) || preorder[*curr] > max {
		return nil
	}

	root := &TreeNode{preorder[*curr], nil, nil}
	*curr++
	root.Left = build1008(preorder, curr, root.Val)
	root.Right = build1008(preorder, curr, max)
	return root
}
