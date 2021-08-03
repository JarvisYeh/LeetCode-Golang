package main

func deleteNode450(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if key < root.Val {
		root.Left = deleteNode450(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode450(root.Right, key)
	} else {
		// now root.Val == key
		// case 1: root has only one child or root has no child
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// case 2: root has two children
		min := root.Right
		for min.Left != nil {
			min = min.Left
		}
		// replace root value with that min value
		root.Val = min.Val
		// recursively delete that min node in right sub tree
		root.Right = deleteNode450(root.Right, min.Val)
	}
	return root
}
