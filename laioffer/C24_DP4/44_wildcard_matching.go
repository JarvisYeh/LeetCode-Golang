package main

func isMatch44(s string, p string) bool {
	DP := make([][]bool, len(s)+1)
	for i, _ := range DP {
		DP[i] = make([]bool, len(p)+1)
	}

	// base cases
	// 1. if s = p ="", true
	DP[0][0] = true
	// 2. s != "", p = "" false
	for i := 1; i <= len(s); i++ {
		DP[i][0] = false
	}
	// 3. s = ""
	// if p[j] = "*" => DP[0][j] = DP[0][j - 1]
	// else DP[0][j] = false
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			DP[0][j] = DP[0][j-1] // p = "***"
		} else {
			DP[0][j] = false
		}
	}

	// induction rule
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			// xxxxxx a
			// yyyyyy ?
			// or
			// xxxxxx a
			// yyyyyy a
			if p[j-1] == '?' || p[j-1] == s[i-1] {
				DP[i][j] = DP[i-1][j-1]
			} else if p[j-1] == '*' {
				// xxxxxx a
				// yyyyyy *
				DP[i][j] = DP[i][j-1] || DP[i-1][j]
			} else {
				// xxxxx a
				// xxxxx b
				DP[i][j] = false
			}
		}
	}
	return DP[len(s)][len(p)]
}
