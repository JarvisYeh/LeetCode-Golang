package main

func isMatch10(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i, _ := range dp {
		dp[i] = make([]bool, n+1)
	}

	// base cases
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		dp[i][0] = false
	}

	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		} else {
			dp[0][j] = false
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// match 0 pre letter
				dp[i][j] = dp[i][j-2]
				// if pre letter is . or pre letter matches
				// can match 1 or more pre letter
				if p[j-2] == '.' || s[i-1] == p[j-2] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[m][n]
}

func main() {
	isMatch10("aa", "a*")
}
