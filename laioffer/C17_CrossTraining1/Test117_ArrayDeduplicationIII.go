package main

import "fmt"

// sorted array, remove duplicates not repeatedly
// 12332 => 122
func dedup(input []int) []int {
	i, j := 0, 0
	for j < len(input) {
		end := j + 1
		for end < len(input) && input[end] == input[j] {
			end++
		}
		// now input[newJ] != input[j], skip those repeat element if exist
		if end-j > 1 { // more than 1 repeat elements
			j = end
		} else { // only 1 element, no repeat elements
			input[i] = input[j]
			i++
			j++
		}
	}
	return input[:i]
}

func main() {
	fmt.Println(dedup([]int{1, 2, 2, 1, 3, 3, 4, 4}))
}
