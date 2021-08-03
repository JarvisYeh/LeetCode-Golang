package main

import "fmt"

// TS : O(n^2)
func strstr85(large, small string) int {
	// corner cases
	if len(small) == 0 {
		return 0
	}
	if len(small) > len(large) {
		return -1
	}
	sArr := []rune(small)
	lArr := []rune(large)
	// 0 1 2 3 4
	// 0 1
	// 5 - 2 = 3
	for i := 0; i <= len(large)-len(small); i++ {
		if checkEquals85(sArr, lArr, i) {
			return i
		}
	}
	return -1
}

func checkEquals85(s, l []rune, i int) bool {
	for j := 0; j < len(s); j++ {
		if s[j] != l[i+j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(strstr85("abcdef", "cd"))
}
