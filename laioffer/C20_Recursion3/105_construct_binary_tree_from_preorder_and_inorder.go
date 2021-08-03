package main

// TC: O(n)
// SC: O(n)
func buildTree(preorder []int, inorder []int) *TreeNode {
	indexMap := make(map[int]int)
	for i, val := range inorder {
		indexMap[val] = i
	}
	return construct(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, indexMap)
}

func construct(pre []int, pL, pR int, in []int, iL, iR int, indexMap map[int]int) *TreeNode {
	if pL > pR {
		return nil
	}
	root := &TreeNode{pre[pL], nil, nil}
	rootIndex := indexMap[pre[pL]]
	lSize := rootIndex - iL
	root.Left = construct(pre, pL+1, pL+lSize, in, iL, rootIndex-1, indexMap)
	root.Right = construct(pre, pL+lSize+1, pR, in, rootIndex+1, iR, indexMap)
	return root
}
