package main

import "fmt"

// return all subsets
// there is duplicate elements in original subset
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

func main() {
	fmt.Println(subSets63("bb"))
}
