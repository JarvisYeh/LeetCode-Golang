package main

import "fmt"

func shiftingLetters848(s string, shifts []int) string {
	bytes := []byte(s)

	shifts[len(shifts)-1] %= 26
	for i := len(shifts) - 1; i >= 1; i-- {
		shifts[i-1] = (shifts[i-1] + shifts[i]) % 26
	}

	for i, val := range bytes {
		bytes[i] = (val-'a'+byte(shifts[i]))%26 + 'a'
	}
	return string(bytes)
}

func main() {
	fmt.Println(shiftingLetters848("benhwjsenjhvulyvefdn", []int{183265101, 732053054, 190062728, 192602923, 551817738, 880105762, 67914564, 336769190, 208588970, 748586819, 57544790, 922070787, 299351164, 42155594, 120233240, 184503545, 976640197, 453293964, 150570427, 82070647}))
}
