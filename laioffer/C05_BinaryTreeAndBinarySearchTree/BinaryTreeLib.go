package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) PrintTree() {
	if root == nil {
		fmt.Println("Empty Tree")
	}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			curr := q[0]
			q = q[1:]
			if curr == nil {
				fmt.Printf("%s\t", "#")
				continue
			}
			fmt.Printf("%d\t", curr.Val)
			if curr.Left != nil {
				q = append(q, curr.Left)
			} else {
				q = append(q, nil)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			} else {
				q = append(q, nil)
			}
		}
		fmt.Println()
	}
}

func FormCompleteTree(arr []string) *TreeNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	nodeArr := make([]*TreeNode, len(arr))
	for i, _ := range nodeArr {
		if val, err := strconv.Atoi(arr[i]); err == nil {
			nodeArr[i] = &TreeNode{val, nil, nil}
		} else {
			panic("can not form a complete tree, check input")
		}
	}
	for i := 0; i < len(nodeArr); i++ {
		if i*2+1 < len(arr) {
			left := nodeArr[2*i+1]
			nodeArr[i].Left = left
		}
		if i*2+2 < len(arr) {
			right := nodeArr[2*i+2]
			nodeArr[i].Right = right
		}
	}
	return nodeArr[0]
}

func FormTreeByLayer(arr []string) *TreeNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	q := []*TreeNode{}
	nodeArr := make([]*TreeNode, len(arr))
	for i, _ := range nodeArr {
		if val, err := strconv.Atoi(arr[i]); err == nil {
			nodeArr[i] = &TreeNode{val, nil, nil}
		}
	}
	if nodeArr[0] == nil {
		panic("invalid input")
	}
	q = append(q, nodeArr[0])

	i := 1
	for len(q) != 0 {
		curr := q[0]
		q = q[1:]

		if curr == nil {
			continue
		}

		if i < len(arr) {
			curr.Left = nodeArr[i]
			q = append(q, nodeArr[i])
			i++
		}

		if i < len(arr) {
			curr.Right = nodeArr[i]
			q = append(q, nodeArr[i])
			i++
		}
	}
	return nodeArr[0]
}

func main() {
	//root1 := FormCompleteTree([]string{"1", "2", "3", "4"})
	root := FormTreeByLayer([]string{"1", "2", "3", "4", "#", "5", "6", "7", "9", "#", "10", "11", "12"})
	//root1.PrintTree()
	root.PrintTree()

}
