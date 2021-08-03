package main

import "fmt"

// method 1:
// for each index i
// look left as far as possible
// look right as for as possible
// TC: O(n^2)
// SC: O(1)
func largestRectangleArea84I(heights []int) int {
	max := 0
	for i := 0; i < len(heights); i++ {
		left := i
		for left >= 0 && heights[left] >= heights[i] {
			left--
		}

		right := i
		for right < len(heights) && heights[right] >= heights[i] {
			right++
		}

		area := heights[i] * (right - left - 1)
		if area > max {
			max = area
		}
	}
	return max
}

// method 2:
// for each index i
// 1. go through array once to find the most left index i could reach
// 2. go through array once to find the most right index i could reach
// the above 2 steps are based on the help of a stack
// for both steps
//	for index i:
//	stack stores the ascending heights index subarray ends at index i - 1
// 	this is for fast looking back
func largestRectangleArea84II(heights []int) int {
	left, right := make([]int, len(heights)), make([]int, len(heights))

	// Stack stores the heights ascending bar's index ends at last explore bar
	stack := []int{}
	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 {
			left[i] = 1
			stack = append(stack, i)
			continue
		}

		for len(stack) != 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		// now stack[end - 1] < heights[i] || len(stack) = 0
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			right[i] = 1
			stack = append(stack, i)
			continue
		}
		for len(stack) != 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		// now stack[end - 1] < heights[i] || len(stack) = 0
		if len(stack) == 0 {
			right[i] = len(heights)
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	max := 0
	for i := 0; i < len(left); i++ {
		area := (right[i] - left[i] - 1) * heights[i]
		if max < area {
			max = area
		}
	}
	return max
}

// method 3
// do not use stack
// jump looking back
func largestRectangleArea84III(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	left[0] = -1
	right[n-1] = n

	for i := 1; i < n; i++ {
		prev := i - 1
		for prev >= 0 && heights[prev] >= heights[i] {
			// update prev
			// since idx = range (left[prev], prev] has heights[idx] >= heights[prev]
			// no need to check them, just skip
			prev = left[prev]
		}
		// now heights[prev] < heights[i] || prev = -1
		left[i] = prev
	}

	for i := n - 2; i >= 0; i-- {
		prev := i + 1
		for prev < n && heights[prev] >= heights[i] {
			// update prev
			// since idx = range [prev, right[prev]) has heights[idx] >= heights[prev]
			// no need to check them, just skip
			prev = right[prev]
		}
		// now heights[prev] < heights[i] || prev = n
		right[i] = prev
	}

	// find max area
	max := 0
	for i := 0; i < len(left); i++ {
		area := (right[i] - left[i] - 1) * heights[i]
		if max < area {
			max = area
		}
	}
	return max
}

func main() {
	fmt.Println(largestRectangleArea84II([]int{2, 1, 5, 6, 2, 3}))
}
