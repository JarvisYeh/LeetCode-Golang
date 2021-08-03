package main

func minCut(s string) int {
	// minCuts[i] represents minimal cuts needed to make [0, i] substring all palindrome partition
	minCuts := make([]int, len(s))
	minCuts[0] = 0

	for i := 1; i < len(s); i++ {
		if isPalindrome(s[0 : i+1]) {
			minCuts[i] = 0
			continue
		}
		min := i // [0, i] needs at most i cuts to make all single letters e.g. palindrome
		// [0, i] => [0, j] U (j, i]
		for j := 0; j < i; j++ {
			if isPalindrome(s[j+1:i+1]) && minCuts[j]+1 < min {
				min = minCuts[j] + 1
			}
		}
		minCuts[i] = min
	}
	return minCuts[len(s)-1]
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return false
}

func main() {
	minCut("aab")
}
