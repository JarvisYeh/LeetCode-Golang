package main

// in place solution
// remove particular characters
func removeVowels1119(s string) string {
	charArr := []rune(s)
	i, j := 0, 0
	for j < len(charArr) {
		if charArr[j] == 'a' || charArr[j] == 'e' || charArr[j] == 'i' || charArr[j] == 'o' || charArr[j] == 'u' {
			j++
		} else {
			charArr[i] = charArr[j]
			i++
			j++
		}
	}
	return string(charArr)
}
