package main

// lca of k-nary tree
func lowestCommonAncestor648(root *KnaryTreeNode, nodes []*KnaryTreeNode) *KnaryTreeNode {
	nodeSet := make(map[*KnaryTreeNode]bool)
	for _, node := range nodes {
		nodeSet[node] = true
	}
	return lca648(root, nodeSet)
}

func lca648(root *KnaryTreeNode, nodesSet map[*KnaryTreeNode]bool) *KnaryTreeNode {
	if root == nil {
		return root
	}
	if _, ok := nodesSet[root]; ok {
		return root
	}

	targetNodes := []*KnaryTreeNode{}
	for _, child := range root.Children {
		if lca648(child, nodesSet) != nil {
			targetNodes = append(targetNodes, child)
		}
	}

	if len(targetNodes) == 0 {
		return nil
	}
	if len(targetNodes) == 1 {
		return targetNodes[0]
	}
	return root
}
