package main

import (
	"fmt"
	"sort"
)

func merge33(intervals [][]int) [][]int {
	// sort the input 2d arrays w.r.t first element
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	newPair := intervals[0]
	for _, pair := range intervals {
		// if the newPair is overlap with the next pair
		if newPair[1] >= pair[0] {
			// update the range of newPair if necessary
			if newPair[1] < pair[1] {
				newPair[1] = pair[1]
			}
		} else {
			res = append(res, newPair)
			newPair = pair
		}
	}
	// append the last pair
	res = append(res, newPair)
	return res
}

func main() {
	fmt.Println(merge33([][]int{{1, 3}, {-2, 2}}))
}
