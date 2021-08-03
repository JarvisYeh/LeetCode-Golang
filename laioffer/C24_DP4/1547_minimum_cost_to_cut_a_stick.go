package main

import (
	"math"
	"sort"
)

func minCost1547(n int, cuts []int) int {
	helper := make([]int, len(cuts)+2)
	helper[0] = 0
	helper[len(helper)-1] = n
	for i := 0; i < len(cuts); i++ {
		helper[i+1] = cuts[i]
	}
	sort.Ints(helper)

	// minCosts[i][j] represents the min cost of cutting wood piece {helper[i], helper[j]}
	minCosts := make([][]int, len(helper))
	for i, _ := range minCosts {
		minCosts[i] = make([]int, len(helper))
	}

	// fill the form from bottom to up, left to right
	// i < j
	for i := len(helper) - 2; i >= 0; i-- {
		for j := i + 1; j < len(helper); j++ {
			if i == j-1 {
				minCosts[i][j] = 0
			} else {
				// cut {helper[i], helper[j]} into {helper[i], helper[t]} and {helper[t], helper[j]}
				min := math.MaxInt64
				for t := i + 1; t < j; t++ {
					if minCosts[i][t]+minCosts[t][j] < min {
						min = minCosts[i][t] + minCosts[t][j]
					}
				}
				// add the cost of cutting {helper[i], helper[j]}
				minCosts[i][j] = min + helper[j] - helper[i]
			}
		}
	}
	return minCosts[0][len(helper)-1]
}
