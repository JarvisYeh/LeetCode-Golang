package main

func stringShift1427(s string, shift [][]int) string {
	r := 0
	for _, val := range shift {
		if val[0] == 0 {
			r -= val[1]
		} else {
			r += val[1]
		}
	}

	// always shift right
	// first mod len(s) to make the abs(r) < len(s)
	// transform shift left to right
	r = r % len(s)
	if r < 0 {
		r = len(s) + r
	}

	arr := []byte(s)
	shiftRight1427(arr, r)
	return string(arr)
}

func shiftRight1427(s []byte, shift int) {
	reverse1427(s, 0, len(s)-1)
	reverse1427(s, 0, shift-1)
	reverse1427(s, shift, len(s)-1)
}

func reverse1427(s []byte, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	stringShift1427("abc", [][]int{{0, 1}, {1, 2}})
}
