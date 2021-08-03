package main

import (
	"fmt"
	"math"
)

func closestKValues504(root *TreeNode, target, k int) []int {
	q := &[]int{}
	inOrder504(root, target, k, q)
	return *q
}

func inOrder504(root *TreeNode, target, k int, q *[]int) {
	if root == nil {
		return
	}

	inOrder504(root.Left, target, k, q)
	if len(*q) < k { // if less than k elements in queue, append that value
		*q = append(*q, root.Val)
	} else if math.Abs(float64(root.Val-target)) < math.Abs(float64((*q)[0]-target)) {
		// if new value if more closer than the oldest element in queue (which is most far away from target)
		// replace that element
		*q = (*q)[1:]
		*q = append(*q, root.Val)
	} else {
		return
	} // if new elements is already further than oldest element in queue, no need to check more elements

	inOrder504(root.Right, target, k, q)
}

func main() {
	root := FormTreeByLayer([]string{"7", "5", "8", "3", "6", "#", "13", "1"})
	root.PrintTree()
	fmt.Println(closestKValues504(root, 6, 3))
}
