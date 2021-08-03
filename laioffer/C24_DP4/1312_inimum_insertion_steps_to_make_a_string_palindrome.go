package main

// TC: O(n^2)
// SC: O(n^2)
func minInsertions1312(s string) int {
	// minSteps[i][j] represents the min steps needed to make s[i, j] a palindrome
	minSteps := make([][]int, len(s))
	for i, _ := range minSteps {
		minSteps[i] = make([]int, len(s))
	}

	// to fill M[i][j], need M[x][y] where x > i, y < j
	// fill larger i first, smaller j first
	// i.e. from left to right, bottom to up
	// and i < j
	for j := 0; j < len(s); j++ {
		for i := j; i >= 0; i-- {
			if i == j {
				minSteps[i][j] = 0
			} else if i == j-1 && s[i] == s[j] {
				minSteps[i][j] = 0
			} else if s[i] == s[j] {
				minSteps[i][j] = minSteps[i+1][j-1]
			} else { // s[i] != s[j] and j - i > 1
				left := minSteps[i][j-1]
				right := minSteps[i+1][j]
				if left > right {
					left = right
				}
				// M[i][j] = 1 + min(left, right)
				minSteps[i][j] = 1 + left
			}
		}
	}
	return minSteps[0][len(s)-1]
}
