package main

func findAndReplacePattern890(words []string, pattern string) []string {
	patArr := []byte(pattern)
	normalize890(patArr)
	res := []string{}

	for _, word := range words {
		if len(word) != len(patArr) {
			continue
		}
		wordArr := []byte(word)
		normalize890(wordArr)
		if compareArr890(wordArr, patArr) {
			res = append(res, word)
		}
	}

	return res
}

// normalize a byte slice
// def -> abc, qqw -> aab
func normalize890(s []byte) {
	// map stores (key = char in s, value = normalized char)
	dict := make(map[byte]byte)
	// start alphabet
	char := byte('a')
	// for each char in slice
	for i, val := range s {
		if _, ok := dict[val]; ok {
			s[i] = dict[val]
		} else {
			dict[val] = char
			s[i] = char
			char++
		}
	}
}

// return true is two byte slice is equal
// false otherwise
func compareArr890(s1, s2 []byte) bool {
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	findAndReplacePattern890([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb")
}
