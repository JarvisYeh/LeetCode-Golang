package main

import (
	"fmt"
	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

// maintain a deque where left -> right always in correct order as tree levels
func zigzagLevelOrder103(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	// create a deque
	dq := deque.New()
	dq.PushRight(root) // for root, left right doesn't matter
	for dq.Size() != 0 {
		// first level print left -> right
		tmp := []int{}
		for i := dq.Size(); i > 0; i-- {
			curr := dq.PopLeft().(*TreeNode)
			tmp = append(tmp, curr.Val)
			// if pop left, definitely push right
			// if push right, check node's left child first
			// to maintain deque left -> right same as tree left -> right
			if curr.Left != nil {
				dq.PushRight(curr.Left)
			}
			if curr.Right != nil {
				dq.PushRight(curr.Right)
			}
		}
		if len(tmp) > 0 {
			res = append(res, tmp)
		}

		// second level print right -> left
		tmp = []int{}
		for i := dq.Size(); i > 0; i-- {
			curr := dq.PopRight().(*TreeNode)
			tmp = append(tmp, curr.Val)
			// if pop right, definitely push left
			// if push left, check node's right child first
			// to maintain deque left -> right same as tree left -> right
			if curr.Right != nil {
				dq.PushLeft(curr.Right)
			}
			if curr.Left != nil {
				dq.PushLeft(curr.Left)
			}
		}
		if len(tmp) > 0 {
			res = append(res, tmp)
		}
	}
	return res
}

func main() {
	root := FormTreeByLayer([]string{"3", "9", "20", "null", "null", "15", "7"})
	fmt.Println(zigzagLevelOrder103(root))
}
