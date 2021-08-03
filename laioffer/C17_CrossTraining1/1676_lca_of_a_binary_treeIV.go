package main

// k nodes' lca of binary tree
// assume k nodes are in the tree
func lowestCommonAncestor1676(root *TreeNode, nodes []*TreeNode) *TreeNode {
	nodeSet := make(map[*TreeNode]bool)
	for _, node := range nodes {
		nodeSet[node] = true
	}
	return lca1676(root, nodeSet)
}

func lca1676(root *TreeNode, nodeSet map[*TreeNode]bool) *TreeNode {
	if root == nil {
		return root
	}
	if _, ok := nodeSet[root]; ok {
		return root
	}

	left := lca1676(root.Left, nodeSet)
	right := lca1676(root.Right, nodeSet)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
