package main

// method 1:
// pre-order traversal
// prev.Right = curr
func flatten114I(root *TreeNode) {
	prev := new(*TreeNode)
	*prev = &TreeNode{-1, nil, nil}
	preOrder114(root, prev)
}

func preOrder114(root *TreeNode, prev **TreeNode) {
	if root == nil {
		return
	}

	left := root.Left
	right := root.Right

	root.Left = nil
	(*prev).Right = root
	*prev = root

	preOrder114(left, prev)
	preOrder114(right, prev)
}

// method 2:
// mirror post-order traversal (right child -> right child -> root)
// curr.Right = curr
func flatten114II(root *TreeNode) {
	prev := new(*TreeNode)
	mirrorPostOrder(root, prev)
}

func mirrorPostOrder(root *TreeNode, prev **TreeNode) {
	if root == nil {
		return
	}

	mirrorPostOrder(root.Right, prev)
	mirrorPostOrder(root.Left, prev)

	// since children are all processed, just deal with current root
	// current level connect with prev node
	root.Left = nil
	root.Right = *prev
	*prev = root
}
