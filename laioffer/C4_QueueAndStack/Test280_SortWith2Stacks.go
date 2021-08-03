package main

import (
	"fmt"
	"math"
)

// use count to record how many elements are unsorted
func sortWith2Stacks280I(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}

	inStack := nums
	bufStack := make([]int, 0)

	// store the amount of unsorted elements
	count := len(nums)
	// store the amount of min in one round
	count_min := 0

	for count > 0 {
		min := inStack[len(inStack)-1]
		// if some unsorted number remains in in stack
		for i := 0; i < count; i++ {
			num := inStack[len(inStack)-1]
			if num < min {
				count_min = 1
				min = num
			} else if num == min {
				count_min++
			}
			inStack = inStack[:len(inStack)-1]
			bufStack = append(bufStack, num)
		}

		// push count_min times of min
		for count_min > 0 {
			inStack = append(inStack, min)
			count_min--
			count--
		}

		// move the unsorted numbers back to inStack
		for len(bufStack) > 0 {
			num := bufStack[len(bufStack)-1]
			if num != min {
				inStack = append(inStack, num)
			}
			bufStack = bufStack[:len(bufStack)-1]
		}
	}
	return inStack
}

// res.peek() >= min as the while loop condition
// res: [1 2 3 | 4 9 12
// 3 < min = 4, not pop 3
func sortWith2Stacks280II(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}
	inStack := make([]int, len(nums))
	copy(inStack, nums)
	resStack := make([]int, 0)

	for len(inStack) > 0 {
		min := math.MaxInt64
		minCount := 0
		// move all the elements from in stack to result stack
		// find the min among them
		for len(inStack) > 0 {
			num := inStack[len(inStack)-1]
			inStack = inStack[:len(inStack)-1]
			if num < min {
				min = num
				minCount = 1
			} else if num == min {
				minCount++
			}
			resStack = append(resStack, num)
		}

		// move the unsorted elements back to inStack
		for len(resStack) > 0 && resStack[len(resStack)-1] >= min {
			num := resStack[len(resStack)-1]
			// pop from buffer
			resStack = resStack[:len(resStack)-1]
			if num != min {
				// push to in
				inStack = append(inStack, num)
			}
		}

		// push the current round min to resStack
		for i := 0; i < minCount; i++ {
			resStack = append(resStack, min)
		}
	}

	return resStack
}

func main() {
	arr := []int{2, 4, 1, 5, 2, 5, 6, 7, 10, 2}
	fmt.Println(sortWith2Stacks280II(arr))
}
