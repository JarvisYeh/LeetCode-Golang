package main

import "fmt"

// TC: O(nlogn)
// SC: O(n)
func findSubsequences682(nums []int) []int {
	lens := make([]int, len(nums))
	lowestEnds := []int{-1}

	maxLen, maxIdx := 1, 0
	for i := 0; i < len(nums); i++ {
		targetIdx := binarySearch682(lowestEnds, nums[i])
		// update lowestEnds array with nums[i]
		// add / replace
		// update lens[i] with nums[i]'s index in lowestEnds
		// index is the longest length of subsequence ends at nums[i]
		if targetIdx == -1 {
			lowestEnds = append(lowestEnds, nums[i])
			lens[i] = len(lowestEnds) - 1
		} else {
			lowestEnds[targetIdx] = nums[i]
			lens[i] = targetIdx
		}

		// update max
		if lens[i] > maxLen {
			maxLen = lens[i]
			maxIdx = i
		}
	}

	// post-processing
	res := make([]int, maxLen)
	for i := len(res) - 1; i >= 0; i-- {
		res[i] = nums[maxIdx]
		j := maxIdx - 1
		for j >= 0 && (lens[j] != lens[maxIdx]-1 || nums[j] >= nums[maxIdx]) {
			j--
		}
		maxIdx = j
	}

	return res
}

func binarySearch682(lowestEnds []int, curr int) int {
	l, r := 1, len(lowestEnds)-1
	for l <= r {
		mid := l + (r-l)/2
		if lowestEnds[mid] < curr {
			l = mid + 1
		} else if mid != 1 && lowestEnds[mid-1] >= curr {
			r = mid - 1
		} else {
			l = mid
			break
		}
	}

	if l <= r {
		return l
	} else {
		return -1
	}
}

func main() {
	fmt.Println(findSubsequences682([]int{28, 4, 8, 14, 14, 12, 7, 14, 28, 24, 9, 30, 28, 29, 26, 3, 17, 18, 5, 29, 18, 8, 30, 32, 13, 29, 6}))
}
