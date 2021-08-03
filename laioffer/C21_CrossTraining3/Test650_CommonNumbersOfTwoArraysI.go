package main

import "fmt"

// unsorted without duplication
// TC : O(m + n)
// SC : O(min(m, n))
func common650(a, b []int) []int {
	res := []int{}
	set := make(map[int]bool)
	// determine smaller size array as a
	if len(b) < len(a) {
		a, b = b, a
	}
	// always put elements in a to set
	for _, num := range a {
		set[num] = true
	}

	for _, num := range b {
		if _, ok := set[num]; ok {
			res = append(res, num)
		}
	}
	return res
}

func main() {
	fmt.Println(common650(
		[]int{1, 9, -5, 6, 7, 2},
		[]int{9, 4, 100, 2}))
}
