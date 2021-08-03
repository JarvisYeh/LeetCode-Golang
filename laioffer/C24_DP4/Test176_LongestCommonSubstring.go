package main

import "fmt"

func longestCommon176(source, target string) string {
	// DP[i][j] represents the longest substring length of
	// first i letters in source and first j letters in target
	// including i-th letter in source and j-th letter in target
	DP := make([][]int, len(source)+1)
	for i, _ := range DP {
		DP[i] = make([]int, len(target)+1)
	}

	start, end, max := 0, 0, 0
	for i := 0; i <= len(source); i++ {
		for j := 0; j <= len(target); j++ {
			if i == 0 || j == 0 {
				DP[i][j] = 0
			} else if source[i-1] == target[j-1] {
				DP[i][j] = DP[i-1][j-1] + 1
				if DP[i][j] > max {
					max = DP[i][j]
					start = i - DP[i][j]
					end = i
				}
			} else {
				DP[i][j] = 0
			}
		}
	}
	return source[start:end]
}

func main() {
	fmt.Println(longestCommon176("abcdefffff", "cdefgfffff"))
}
