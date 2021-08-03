package main

import "fmt"

// determine index by index
// use swap
func keepDistance264(k int) []int {
	res, used := new([]int), new([]bool)
	*res, *used = make([]int, 2*k), make([]bool, k+1)
	for i := 0; i < k; i++ {
		(*res)[i] = i + 1
		(*res)[i+k] = i + 1
	}
	if dfs264(0, res, used, k) {
		return *res
	} else {
		return nil
	}
}

func dfs264(idx int, res *[]int, used *[]bool, k int) bool {
	// base case: all index are placed with number
	if idx == len(*res) {
		return true
	}

	// determine res[idx]
	for i := idx; i < len(*res); i++ {
		num := (*res)[i]
		// 1. num first time occurs
		if !(*used)[num] {
			(*res)[i], (*res)[idx] = (*res)[idx], (*res)[i]
			(*used)[num] = true
			if dfs264(idx+1, res, used, k) {
				return true
			}
			(*used)[num] = false
			(*res)[i], (*res)[idx] = (*res)[idx], (*res)[i]
		} else if idx-num-1 >= 0 && (*res)[idx-num-1] == num {
			// 2. num second time occurs, but res[idx - num - 1] == num
			(*res)[i], (*res)[idx] = (*res)[idx], (*res)[i]
			if dfs264(idx+1, res, used, k) {
				return true
			}
			(*res)[i], (*res)[idx] = (*res)[idx], (*res)[i]
		}
	}
	return false
}

func main() {
	fmt.Println(keepDistance264(4))
}
