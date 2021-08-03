package main

func upsideDownBinaryTree156(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	newRoot := upsideDownBinaryTree156(root.Left)
	root.Left.Left = root.Right
	root.Left.Right = root
	root.Left = nil
	root.Right = nil
	return newRoot
}

func main() {
	root := FormTreeByLayer([]string{"1", "2", "3", "4", "5", "#", "#", "7", "8", "#", "#", "9", "10"})
	root.PrintTree()
	root = upsideDownBinaryTree156(root)
	root.PrintTree()
}
