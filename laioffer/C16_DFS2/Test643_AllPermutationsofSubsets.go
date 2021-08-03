package main

import "fmt"

func allPermutationsOfSubsets643(s string) []string {
	res := &[]string{}
	dfs643(0, []byte(s), res)
	return *res
}

func dfs643(index int, arr []byte, res *[]string) {
	*res = append(*res, string(arr[:index]))
	if index == len(arr) {
		return
	}

	for i := index; i < len(arr); i++ {
		arr[i], arr[index] = arr[index], arr[i]
		dfs643(index+1, arr, res)
		arr[i], arr[index] = arr[index], arr[i]
	}
}

func main() {
	fmt.Println(allPermutationsOfSubsets643("abc"))
}
