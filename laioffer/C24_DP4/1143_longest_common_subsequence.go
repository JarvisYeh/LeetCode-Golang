package main

func longestCommonSubsequence1143(text1 string, text2 string) int {
	DP := make([][]int, len(text1)+1)
	for i, _ := range DP {
		DP[i] = make([]int, len(text2)+1)
	}

	for i := 0; i <= len(text1); i++ {
		for j := 0; j <= len(text2); j++ {
			if i == 0 || j == 0 {
				DP[i][j] = 0
			} else if text1[i-1] != text2[j-1] {
				DP[i][j] = DP[i-1][j-1]
				if DP[i-1][j] > DP[i][j] {
					DP[i][j] = DP[i-1][j]
				}
				if DP[i][j-1] > DP[i][j] {
					DP[i][j] = DP[i][j-1]
				}
			} else {
				DP[i][j] = DP[i-1][j-1] + 1
			}
		}
	}
	return DP[len(text1)][len(text2)]
}
