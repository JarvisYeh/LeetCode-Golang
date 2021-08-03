package main

import "fmt"

// sorted k array with duplication, find common elements

// binary reduction
// TC: O(kn)
// SC: O(n + log k)
func commonElementsInKSortedArrays644(arrs [][]int) []int {
	if arrs == nil || len(arrs) == 0 {
		return []int{}
	}
	if len(arrs) == 1 {
		return arrs[0]
	}
	mid := len(arrs) / 2
	left := commonElementsInKSortedArrays644(arrs[0:mid])
	right := commonElementsInKSortedArrays644(arrs[mid:len(arrs)])
	return merge644(left, right)
}

func merge644(left, right []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			i++
		} else if left[i] > right[j] {
			j++
		} else {
			res = append(res, left[i])
			i++
			j++
		}
	}
	return res
}

func main() {
	fmt.Println(commonElementsInKSortedArrays644([][]int{
		{1, 2, 2, 3, 9},
		{1, 2, 4, 5, 9},
		{1, 2, 5, 7, 9},
		{1, 9, 10},
	}))
}
