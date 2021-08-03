package main

import "math"

func closestValue(root *TreeNode, target float64) int {
	res := root.Val
	for root != nil {
		// check current node
		if math.Abs(float64(root.Val)-target) < math.Abs(float64(res)-target) {
			res = root.Val
		}

		// check which branch to go
		if float64(root.Val) < target {
			root = root.Right
		} else if float64(root.Val) > target {
			root = root.Left
		} else {
			return root.Val
		}
	}
	return res
}
