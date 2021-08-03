package main

func reverseStr541(s string, k int) string {
	arr := []byte(s)
	// j always points to the start location
	for j := 0; j < len(arr); j += 2 * k {
		if j+k > len(arr) {
			reverse541(arr, j, len(arr)-1)
		} else {
			reverse541(arr, j, j+k-1)
		}
	}
	return string(arr)
}

func reverse541(s []byte, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
