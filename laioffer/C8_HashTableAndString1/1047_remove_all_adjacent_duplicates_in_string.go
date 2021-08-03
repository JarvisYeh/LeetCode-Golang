package main

// continually remove two adjacent characters
// aaa => a
// baaabb => ab
func removeDuplicates1047(s string) string {
	arr := []rune(s)
	i, j := 0, 0
	for j < len(arr) {
		if i == 0 || arr[j] != arr[i-1] {
			arr[i] = arr[j]
			i++
			j++
		} else { // arr[j] == arr[i - 1]
			j++
			i--
		}
	}
	return string(arr[:i])
}
