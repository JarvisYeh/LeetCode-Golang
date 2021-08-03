package main

func romanToInt13(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	i := 0
	for i < len(s) {
		curr := dict[s[i]]
		// if next > curr
		// e.g. IV = 5 - 1 = 4
		if i+1 < len(s) && dict[s[i+1]] > curr {
			sum += dict[s[i+1]] - curr
			i += 2
		} else {
			sum += curr
			i++
		}
	}
	return sum

}
