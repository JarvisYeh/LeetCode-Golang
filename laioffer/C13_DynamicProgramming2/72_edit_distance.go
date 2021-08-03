package main

import "fmt"

// DFS
// TC: O(3^(m+n))
// 3 child for each nodes
// recursion tree has at most (m + n) levels (delete s1 and s2 alternately)
// SC: O(m + n)
func minDistance72I(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	if word1[0] == word2[0] {
		// if first letter is same, same as minDistance(s1[1:], s2[1:])
		return minDistance72I(word1[1:], word2[1:])
	} else {
		// option 1: replace s1[0]
		replace := 1 + minDistance72I(word1[1:], word2[1:])
		// option 2: delete s1[0]
		del := 1 + minDistance72I(word1[1:], word2)
		// option 3: insert s2[0] before s1[0]
		// same as delete s2[0]
		insert := 1 + minDistance72I(word1, word2[1:])
		min := replace
		if del < min {
			min = del
		}
		if insert < min {
			min = insert
		}
		return min
	}
}

// DP
// TC : O(m*n)
// SC: O(m*n)
func minDistance72II(word1, word2 string) int {
	m, n := len(word1), len(word2)
	// distance[i][j] represents the minimal distance
	// of first i letters in s1 and first j letters in s2
	distance := make([][]int, m+1)
	for i, _ := range distance {
		distance[i] = make([]int, n+1)
	}

	// base case
	// s1 = "" -> s2[:j + 1]
	// s1[:i + 1] -> s2 = ""
	for i := 0; i <= m; i++ {
		distance[i][0] = i
	}
	for j := 0; j <= n; j++ {
		distance[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				distance[i][j] = distance[i-1][j-1]
			} else {
				// option 1: replace s1[i - 1] with s2[j - 1]
				replace := 1 + distance[i-1][j-1]
				// option 2: delete s1[i - 1]
				del := 1 + distance[i-1][j]
				// option 3: insert s2[j - 1] after s1[i - 1]
				// same as delete s2[j - 1]
				insert := 1 + distance[i][j-1]

				min := replace
				if del < min {
					min = del
				}
				if insert < min {
					min = insert
				}
				distance[i][j] = min
			}
		}
	}
	return distance[m][n]
}

func main() {
	fmt.Println(minDistance72II("dinitrophenylhydrazine", "acetylphenylhydrazine"))
}
