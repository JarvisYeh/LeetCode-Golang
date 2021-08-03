package main

import "fmt"

// sorted array
// use two pointer
// TC: O(n)
// SC: O(1)
func existI180(nums []int, target int) bool {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			return true
		}
	}
	return false
}

// unsorted array
// use set to store nums[0:i]
// currently check whether (target - nums[i]) is in set
// TC: O(n)
// SC: O(n)
func existII180(nums []int, target int) bool {
	set := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		// if target - nums[i] exist in nums[0:i], exist result
		if _, ok := set[target-nums[i]]; ok {
			return true
		}
		set[nums[i]] = true
	}
	return false
}

// unsorted array
// sort first, then use two pointer
// TC: O(nlogn)
// SC: O(1)
func existIII180(nums []int, target int) bool {
	heapSort(nums)
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			return true
		}
	}
	return false
}

func heapSort(nums []int) {
	// heapify O(n)
	// start from the last parent index
	for i := len(nums)/2 - 1; i >= 0; i-- {
		candidate := i
		if nums[i*2+1] > candidate {
			candidate = i*2 + 1
		}
		if nums[i*2+2] > candidate {
			candidate = i*2 + 2
		}
		// swap larger to parent node
		nums[i], nums[candidate] = nums[candidate], nums[i]
	}

	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		// swap current top to end, same as remove from heap
		nums[0], nums[i] = nums[i], nums[0]
		// percolate down until i - 1, maintain heap property
		j := 0
		for j < i {
			candidate := j
			if j*2+1 < i && nums[candidate] < nums[j*2+1] {
				candidate = j*2 + 1
			}
			if j*2+2 < i && nums[candidate] < nums[j*2+2] {
				candidate = j*2 + 2
			}

			if j == candidate {
				break
			} else {
				nums[j], nums[candidate] = nums[candidate], nums[j]
				j = candidate
			}
		}
	}
}

func main() {
	fmt.Println(existIII180([]int{3, 5, 2, 2, 10, 200, -6}, -6))
}
