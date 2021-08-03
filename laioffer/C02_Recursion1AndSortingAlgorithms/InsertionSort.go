package main

import "fmt"

func insertionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > key; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = key
	}
	return nums
}

func main() {
	fmt.Println(insertionSort([]int{1, 1, 10, 4}))
}
