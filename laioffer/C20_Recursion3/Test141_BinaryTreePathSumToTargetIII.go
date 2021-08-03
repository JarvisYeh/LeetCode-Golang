package main

// whether there exist a sub-path from root to leaf whose sum = target

// method 1:
// use a preNodes array stores all node value from root to curr
// TC: O(n*h) every node has a for loop with size of tree height
// SC: O(h)
func exist141I(root *TreeNode, target int) bool {
	return existHelperI141(root, []int{}, target)
}

// preNodes[i] stores the i-th node value from root downward
func existHelperI141(root *TreeNode, preNodes []int, target int) bool {
	// corner case
	if root == nil {
		return false
	}

	// current level check upwards if there exist a sub-path (including current node), sum to target
	sum := root.Val
	if sum == target {
		return true
	}
	for i := len(preNodes) - 1; i >= 0; i-- {
		sum += preNodes[i]
		if sum == target {
			return true
		}
	}

	left, right := false, false
	if root.Left != nil {
		left = existHelperI141(root.Left, append(preNodes, root.Val), target)
	}
	if root.Right != nil {
		right = existHelperI141(root.Right, append(preNodes, root.Val), target)
	}
	return left || right
}

// method 2:
// use a prefixSums set
// stores sum of nodes from root to i-the node downwards i = [0, curr)
// TC : O(n), check set needs O(1)
// SC : O(h)
func exist141II(root *TreeNode, target int) bool {
	set := make(map[int]bool)
	set[0] = true
	return helper141II(root, 0, set, target)
}

func helper141II(root *TreeNode, sum int, set map[int]bool, target int) bool {
	if root == nil {
		return false
	}

	// current level check if (sum - target) in set
	sum += root.Val
	if _, ok := set[sum-target]; ok {
		return true
	}

	// add curr sum to set
	_, exist := false, false
	if _, exist = set[sum]; !exist {
		set[sum] = true
	}

	// ask for children
	left := helper141II(root.Left, sum, set, target)
	right := helper141II(root.Right, sum, set, target)

	// if add curr sum in this level, remove and return to previous level
	if !exist {
		delete(set, sum)
	}
	return left || right
}
