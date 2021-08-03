package main

func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]bool)
	max := 1

	left, right := 0, 0 // window = [left, right)
	for right < len(s) {
		if set[s[right]] {
			delete(set, s[left])
			left++
		} else {
			set[s[right]] = true
			right++
			if max < right-left {
				max = right - left
			}
		}
	}
	return max

}
