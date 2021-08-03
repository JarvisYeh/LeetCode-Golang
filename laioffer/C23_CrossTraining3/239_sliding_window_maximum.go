package main

import "fmt"

type Pair struct {
	Val int
	Idx int
}

// TC : O(n) every element push/pop at most once
// SC: O(k)
func maxSlidingWindow239(nums []int, k int) []int {
	deque := []Pair{}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		// pop until deque[right] > nums[i]
		j, n := 0, len(deque)
		for n-1-j >= 0 && deque[n-1-j].Val < nums[i] {
			j++
		}
		deque = deque[0 : n-j]
		// offer current element
		deque = append(deque, Pair{nums[i], i})

		// if first window not filled, skip finding max
		if i < k-1 {
			continue
		}
		// check if deque[0] is in current window range
		if deque[0].Idx < i-k+1 || deque[0].Idx > i {
			deque = deque[1:]
		}
		res = append(res, deque[0].Val)
	}
	return res
}

func main() {
	fmt.Println(maxSlidingWindow239([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
