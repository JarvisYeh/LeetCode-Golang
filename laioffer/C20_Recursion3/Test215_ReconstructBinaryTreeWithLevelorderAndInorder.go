package main

// TC: O(n^2)
// SC: O(n^2)
func reconstruct(level, in []int) *TreeNode {
	inMap := make(map[int]int)
	for i, val := range in {
		inMap[val] = i
	}
	return build(level, inMap)
}

func build(level []int, inMap map[int]int) *TreeNode {
	if len(level) == 0 {
		return nil
	}

	root := &TreeNode{level[0], nil, nil}
	left := []int{}
	right := []int{}
	// check inorder index map for level[1:end)
	// if its index < root.index in inorder array, it's in left subtree
	// else in right subtree
	for i := 1; i < len(level); i++ {
		if inMap[level[i]] < inMap[level[0]] {
			left = append(left, level[i])
		} else {
			right = append(right, level[i])
		}
	}
	root.Left = build(left, inMap)
	root.Right = build(right, inMap)
	return root
}

func main() {
	root := reconstruct([]int{20, 8, 22, 4, 12, 10, 14}, []int{4, 8, 10, 12, 14, 20, 22})
	root.PrintTree()
}
