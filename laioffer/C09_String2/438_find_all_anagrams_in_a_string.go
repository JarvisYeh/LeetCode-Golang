package main

import "fmt"

// TC: O(s1 + s2)
// SC: O(s2)
func findAnagrams438(s string, p string) []int {
	res := []int{}
	if len(s) < len(p) {
		return res
	}
	diffMap := make(map[byte]int)
	match := 0
	// update diffMap with target string
	for _, c := range p {
		if count, ok := diffMap[byte(c)]; ok {
			diffMap[byte(c)] = count + 1
		} else {
			diffMap[byte(c)] = 1
		}
	}

	// update diffMap with first sliding window
	for i := 0; i < len(p); i++ {
		c := byte(s[i])
		if count, ok := diffMap[c]; ok {
			if count == 0 {
				match--
			}
			count--
			diffMap[c] = count
			if count == 0 {
				match++
			}
		}
	}

	// check first sliding window
	if match == len(diffMap) {
		res = append(res, 0)
	}

	// move sliding window one step a time
	// i is the start of the new sliding window
	for i := 1; i <= len(s)-len(p); i++ {
		leftC, rightC := byte(s[i-1]), byte(s[i+len(p)-1])
		// leftC left window, count++
		if count, ok := diffMap[leftC]; ok {
			if count == 0 {
				match--
			}
			count++
			diffMap[leftC] = count
			if count == 0 {
				match++
			}
		}

		// rightC enter window, count++
		if count, ok := diffMap[rightC]; ok {
			if count == 0 {
				match--
			}
			count--
			diffMap[rightC] = count
			if count == 0 {
				match++
			}
		}
		// if match amount = size of map, find one result
		if match == len(diffMap) {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	fmt.Println(findAnagrams438("abacbabc", "abc"))
}
