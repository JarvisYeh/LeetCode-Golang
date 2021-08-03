package main

import "fmt"

func bubbleSort(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j <= i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j] // swap j and j + 1
			}
		}
	}
	return nums
}

func main() {
	fmt.Println(bubbleSort([]int{3, 1, 2, 7, 10, 2}))
}
