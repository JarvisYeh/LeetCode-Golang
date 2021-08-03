package main

import "fmt"

func getLengthOfOptimalCompression(s string, k int) int {
	chars := []byte{}
	min := len(s)
	dfs(s, &chars, 0, k, &min)
	return min
}

func dfs(s string, chars *[]byte, index, remain int, min *int) {
	// base case
	if index >= len(s) {
		tmp := make([]byte, len(*chars))
		copy(tmp, *chars)
		size := compress(tmp)
		if size < *min {
			*min = size
		}
		return
	}

	if remain > 0 {
		dfs(s, chars, index+1, remain-1, min)
	}
	*chars = append(*chars, s[index])
	dfs(s, chars, index+1, remain, min)
	*chars = (*chars)[:len(*chars)-1]
}

func compress(chars []byte) int {
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
			reverse(chars, start, i-1)
		}
		// update j
		j = k
	}
	return i
}

func reverse(chars []byte, i, j int) {
	for i < j {
		chars[i], chars[j] = chars[j], chars[i]
		i++
		j--
	}
}

func main() {
	fmt.Println(getLengthOfOptimalCompression("abcdefghijklmnopqrstuvwxyz", 16))
}
