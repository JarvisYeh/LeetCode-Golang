package main

import "fmt"

// return all subsets
// there is duplicate elements in original subset

// method 1:
// recursion tree
// each level represents an index in input array
// two branch for each node: add/not add
func subSets63(input string) []string {
	res := &[]string{}
	dfs63(0, []byte{}, input, res)
	return *res
}

func dfs63(index int, curr []byte, input string, res *[]string) {
	if index == len(input) {
		*res = append(*res, string(curr))
		return
	}

	// add input[index]
	dfs63(index+1, append(curr, input[index]), input, res)

	// not add input[index]
	// then not add all those input[index'] where input[index'] = input[index]
	for index+1 < len(input) && input[index+1] == input[index] {
		index++
	}
	// now index == len(input) - 1 || input[index + 1] != input[index]
	dfs63(index+1, curr, input, res)
}

// method 2:
// count number of each letter
// recursion tree
// each level represents a letter
// each branch represents how many letter used for a letter
func subsetsWithDup90(nums []int) [][]int {
	res := [][]int{}
	// frequency map of nums
	countMap := make(map[int]int)
	// keySet store as slice
	keySet := []int{}
	// count each key frequency
	for _, num := range nums {
		if count, ok := countMap[num]; ok {
			countMap[num] = count + 1
		} else {
			countMap[num] = 1
			keySet = append(keySet, num)
		}
	}

	helper90(0, keySet, countMap, []int{}, &res)
	return res
}

func helper90(idx int, keySet []int, countMap map[int]int, curr []int, res *[][]int) {
	if idx == len(keySet) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}

	oldLen := len(curr)
	for i := 0; i <= countMap[keySet[idx]]; i++ {
		for j := 0; j < i; j++ {
			curr = append(curr, keySet[idx])
		}
		helper90(idx+1, keySet, countMap, curr, res)
		curr = curr[:oldLen]
	}
}

func main() {
	fmt.Println(subSets63("bb"))
}
