package main

func pathSum113(root *TreeNode, targetSum int) [][]int {
	res := &[][]int{}
	helper113(root, targetSum, []int{}, res)
	return *res
}

// curr is prefix list
func helper113(root *TreeNode, targetSum int, curr []int, res *[][]int) {
	// corner case
	if root == nil {
		return
	}

	// base case, leaf node
	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			tmp = append(tmp, root.Val)
			*res = append(*res, tmp)
		}
		return
	}

	helper113(root.Left, targetSum-root.Val, append(curr, root.Val), res)

	helper113(root.Right, targetSum-root.Val, append(curr, root.Val), res)

}

func main() {
	root := FormTreeByLayer([]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "5", "1"})
	pathSum113(root, 22)
}
