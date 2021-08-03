package main

func wordBreak139(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, str := range wordDict {
		dict[str] = true
	}

	arr := []byte(s)
	// can break contains if arr[:i] can be break into dict words
	canBreak := make([]bool, len(arr)+1)
	canBreak[0] = true
	for i := 1; i <= len(arr); i++ {
		can := false
		for j := 0; j < i; j++ {
			left := canBreak[j]
			right, _ := dict[string(arr[j:i])]
			if left && right {
				can = true
				break
			}
		}
		canBreak[i] = can
	}
	return canBreak[len(arr)]

}
