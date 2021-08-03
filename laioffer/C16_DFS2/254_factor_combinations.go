package main

import (
	"sort"
)

func getFactors254(n int) [][]int {
	factors := factors(n)
	if len(factors) == 0 {
		return [][]int{}
	}

	// sort from large to small
	sort.Slice(factors, func(i, j int) bool {
		return factors[i] > factors[j]
	})
	res := &[][]int{}
	getFactorComb254(0, n, factors, []int{}, res)
	return *res
}

func getFactorComb254(index, remain int, factors, curr []int, res *[][]int) {
	// base case, no factors remain, and n == 1
	if index == len(factors) {
		if remain == 1 {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			*res = append(*res, tmp)
		}
		return
	}

	// do not add factors[index]
	getFactorComb254(index+1, remain, factors, curr, res)

	// add i factors into curr, i from 1 to until remain % factors[index] != 0
	size := len(curr)
	for remain%factors[index] == 0 {
		curr = append(curr, factors[index])
		remain /= factors[index]
		getFactorComb254(index+1, remain, factors, curr, res)
	}
	curr = curr[:size]
}

func factors(n int) []int {
	factors := []int{}
	for i := n / 2; i > 1; i-- {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func main() {
	getFactors254(12)
}
