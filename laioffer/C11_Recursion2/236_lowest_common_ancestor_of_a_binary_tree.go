package main

func LCA236(root, p, q *TreeNode) *TreeNode {
	// if res == nil, p and q are not in tree root
	res := lowestCommonAncestor(root, p, q)
	if res == nil {
		return nil
	}
	// if res == c, p and q has LCA = c
	if res != p && res != q {
		return res
	}

	// if res == p || res == q
	// 1. res == p, try to find q in subtree root at p
	// if not found, p and q are not in the same tree
	if res == p {
		if lowestCommonAncestor(p, q, q) == q {
			return p
		} else {
			return nil
		}
	}
	// 2. res == q, try to find p in subtree root at q
	// if not found, p and q are not in the same tree
	if res == q {
		if lowestCommonAncestor(q, p, p) == p {
			return q
		} else {
			return nil
		}
	}
	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil { // if both left and right are not null, return root
		return root
	} else if left != nil { // if only left is not null, return left
		return left
	} else { // if only right is not null or both are null, return right
		return right
	}
}
