package main

import (
	"fmt"
)

func largestAndSmallest119(input []int) []int {
	n := len(input)
	// O(n/2) compare, swap loser to right half
	for i := 0; i < n/2; i++ {
		if input[i] < input[n-1-i] {
			input[i], input[n-1-i] = input[n-1-i], input[i]
		}
	}

	max, min := input[0], input[n/2]
	for i := 0; i <= n/2; i++ { //O(n/2)
		if max < input[i] {
			max = input[i]
		}
	}

	for i := n / 2; i < n; i++ { //O(n/2)
		if min > input[i] {
			min = input[i]
		}
	}
	return []int{max, min}
}

func main() {
	fmt.Println(largestAndSmallest119([]int{1, 2, 3, 4, 5, 6, 100, -5}))
}
