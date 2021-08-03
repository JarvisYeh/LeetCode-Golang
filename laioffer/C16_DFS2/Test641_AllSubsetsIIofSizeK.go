package main

import (
	"fmt"
	"sort"
)

func subSetsIIOfSizeK641(input string, k int) []string {
	res := &[]string{}
	arr := []byte(input)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	dfs641(0, []byte{}, arr, k, res)
	return *res
}

func dfs641(index int, curr, arr []byte, k int, res *[]string) {
	if len(curr) == k {
		*res = append(*res, string(curr))
		return
	}

	if index == len(arr) {
		return
	}

	for i := index; i < len(arr); i++ {
		if i == index || arr[i] != arr[i-1] {
			dfs641(i+1, append(curr, arr[i]), arr, k, res)
		}
	}
	dfs641(len(arr), curr, arr, k, res)
}

func main() {
	fmt.Println(subSetsIIOfSizeK641("abb", 2))
}
