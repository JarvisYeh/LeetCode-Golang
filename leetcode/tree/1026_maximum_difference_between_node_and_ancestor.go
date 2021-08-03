package main

import "math"

// find a max and a min in one up-down straight path
func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDiff := new(int)
	*maxDiff = math.MinInt64
	findMax(root, root.Val, root.Val, maxDiff)
	return *maxDiff
}

func findMax(root *TreeNode, min, max int, maxDiff *int) {
	if root == nil {
		if max-min > *maxDiff {
			*maxDiff = max - min
		}
		return
	}

	if root.Val < min {
		min = root.Val
	}

	if root.Val > max {
		max = root.Val
	}
	findMax(root.Left, min, max, maxDiff)
	findMax(root.Right, min, max, maxDiff)
}

func main() {
	root := FormTreeByLayer([]string{"8", "3", "10", "1", "6", "null", "14", "null", "null", "4", "7", "13"})
	maxAncestorDiff(root)
}
