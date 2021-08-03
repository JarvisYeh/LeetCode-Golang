package main

import "fmt"

// unsorted array, remove duplicates repeatedly
func dedup118(input []int) []int {
	i, j := 0, 0
	for j < len(input) {
		if i == 0 || input[j] != input[i-1] {
			input[i] = input[j]
			i++
			j++
		} else {
			for j < len(input) && input[j] == input[i-1] {
				j++
			}
			i--
		}
	}
	return input[:i]
}

func main() {
	fmt.Println(dedup([]int{1, 2, 2, 1, 3, 3, 4, 4}))
}
