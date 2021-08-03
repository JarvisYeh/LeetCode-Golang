package main

import (
	"fmt"
)

func compress443(chars []byte) int {
	i, j := 0, 0
	for j < len(chars) {
		k := j
		for k < len(chars) && chars[k] == chars[j] {
			k++
		}
		// for single character
		if k-j == 1 {
			chars[i] = chars[j]
			i++
		} else { // for continuous character
			count := k - j
			chars[i] = chars[j]
			i++
			start := i
			for count != 0 {
				chars[i] = byte(int('0') + count%10)
				count /= 10
				i++
			}
			reverse443(chars, start, i-1)
		}
		// update j
		j = k
	}
	return i
}

func reverse443(chars []byte, i, j int) {
	for i < j {
		chars[i], chars[j] = chars[j], chars[i]
		i++
		j--
	}
}

func main() {
	input := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	k := compress443(input)
	fmt.Println(string(input[:k]))
}
