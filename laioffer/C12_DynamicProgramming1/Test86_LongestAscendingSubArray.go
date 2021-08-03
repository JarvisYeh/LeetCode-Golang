package main

import "fmt"

func longest86(array []int) int {
	if array == nil || len(array) == 0 {
		return -1
	}
	size, max := 1, array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			size++
			if size > max {
				max = size
			}
		} else {
			size = 1
		}
	}
	return max
}

func main() {
	fmt.Println(longest86([]int{1, 0, 2, 9, 10, 12, 15, 17, 19, -1}))
}
