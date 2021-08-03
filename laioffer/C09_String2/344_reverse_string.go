package main

func reverseString344Iter(s []byte) {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseString344Recur(s []byte, i, j int) {
	if i >= j {
		return
	}

	s[i], s[j] = s[j], s[i]
	reverseString344Recur(s, i+1, j-1)
}
