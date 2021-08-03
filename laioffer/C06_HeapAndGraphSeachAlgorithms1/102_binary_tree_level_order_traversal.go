package main

// Use BFS
func levelOrder57BFS(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		var list []int
		for i := len(q); i > 0; i-- {
			// popped one node
			curr := q[0]
			q = q[1:len(q)]
			list = append(list, curr.Val)
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		res = append(res, list)
	}
	return res
}

// Use DFS
func levelOrder57DFS(root *TreeNode) [][]int {
	res := &[][]int{}
	levelOrder57DFSRecur(root, 0, res)
	return *res
}

func levelOrder57DFSRecur(root *TreeNode, level int, res *[][]int) {
	// base case
	if root == nil {
		return
	}

	// if current level does not has a list for it, create one empty list and append to the res list
	if level+1 > len(*res) {
		*res = append(*res, []int{})
	}

	// add current node into the corresponding list
	(*res)[level] = append((*res)[level], root.Val)

	// recursively call children
	levelOrder57DFSRecur(root.Left, level+1, res)
	levelOrder57DFSRecur(root.Right, level+1, res)
}
