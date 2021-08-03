package main

func firstUniqChar387(s string) int {
	freq := make([]int, 26)
	index := make([]int, 26)
	chars := []byte(s)
	for i, val := range chars {
		j := int(val - 'a')
		if index[j] == 0 {
			index[j] = i
		}
		freq[j]++
	}

	min := len(s)
	for i, val := range freq {
		if val == 1 && index[i] < min {
			min = index[i]
		}
	}
	if min == len(s) {
		return -1
	} else {
		return min
	}
}
