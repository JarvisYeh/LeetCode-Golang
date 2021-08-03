package main

import "fmt"

// return subset with size = k
// there is no duplicate elements in original subset
func subSetsOfSizeK640(input string, k int) []string {
	res := &[]string{}
	dfs640(0, []byte{}, input, k, res)
	return *res
}

func dfs640(index int, curr []byte, input string, k int, res *[]string) {
	// base case 1: all ready added k elements
	if len(curr) == k {
		*res = append(*res, string(curr))
		return
	}

	// base case 2: index reaches end
	if index == len(input) {
		return
	}

	dfs640(index+1, append(curr, input[index]), input, k, res)
	dfs640(index+1, curr, input, k, res)
}

func main() {
	fmt.Println(subSetsOfSizeK640("abcd", 2))
}
