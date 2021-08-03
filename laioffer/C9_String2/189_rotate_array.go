package main

import "fmt"

func rotate189(nums []int, k int) {
	k = k % len(nums)
	reverse189(nums, 0, len(nums)-1)
	reverse189(nums, 0, k-1)
	reverse189(nums, k, len(nums)-1)
}

func reverse189(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	input := []int{1, 2, 3, 4}
	fmt.Println(input)
	rotate189(input, 2)
	fmt.Println(input)
}
