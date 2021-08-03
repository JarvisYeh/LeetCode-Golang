package main

// construct tree from int array
// find max num index i
// [0: i) left subtree
// (i, end] right subtree
func constructMaximumBinaryTree(nums []int) *TreeNode {
	// base case
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}

	mIndex := findMaxIndex(nums)
	root := &TreeNode{nums[mIndex], nil, nil}
	root.Left = constructMaximumBinaryTree(nums[:mIndex])
	root.Right = constructMaximumBinaryTree(nums[mIndex+1:])
	return root
}

func findMaxIndex(nums []int) int {
	max := nums[0]
	index := 0
	for i, v := range nums {
		if v > max {
			max = v
			index = i
		}
	}
	return index
}
