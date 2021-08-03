package main

type KnaryTreeNode struct {
	Left     *KnaryTreeNode
	Right    *KnaryTreeNode
	Val      int
	Children []*KnaryTreeNode
}

// lca of k-nary tree
func lowestCommonAncestor647(root, a, b *KnaryTreeNode) *KnaryTreeNode {
	if root == nil || root == a || root == b {
		return root
	}

	targetChildrens := []*KnaryTreeNode{}
	for _, child := range root.Children {
		if child == a || child == b {
			targetChildrens = append(targetChildrens, child)
		}
	}

	if len(targetChildrens) == 0 {
		return nil
	}
	if len(targetChildrens) == 1 {
		return targetChildrens[0]
	}
	return root
}
