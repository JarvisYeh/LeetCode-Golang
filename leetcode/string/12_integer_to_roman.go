package main

import (
	"fmt"
	"math"
	"strconv"
)

func intToRoman12(num int) string {
	dict := map[int]byte{
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}

	// 99 -> '9', '9'
	digits := []rune(strconv.Itoa(num))
	res := []byte{}
	n := len(digits)
	for i, digit := range digits {
		digit = digit - '0'
		if digit == 9 {
			// 90 = 100 - 10 => C - X = XC
			// append 10 first, then 100
			tmp := int(math.Pow10(n - i - 1))
			res = append(res, dict[tmp])
			res = append(res, dict[10*tmp])
		} else if digit == 4 {
			// 40 = 50 - 10 => L - X = XL
			// append 10 first, then 50
			tmp := int(math.Pow10(n - i - 1))
			res = append(res, dict[tmp])
			res = append(res, dict[5*tmp])
		} else {
			// 70 = 50 + 10 + 10 => L + X + X = LXX
			// append 50 first, then 10, then 10
			tmp := int(math.Pow10(n - i - 1))
			if digit/5 != 0 {
				res = append(res, dict[5*tmp])
				digit %= 5
			}
			for digit != 0 {
				res = append(res, dict[tmp])
				digit--
			}
		}
	}
	return string(res)
}
func main() {
	fmt.Println(intToRoman12(58))
}
