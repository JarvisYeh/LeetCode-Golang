package main

// TC : O(n^2)
// SC : O(n)
func maxProduct318(words []string) int {
	intDict := make([]int, len(words))
	// pre-processing
	// intDict[i]'s i-th bit = 1
	// means in words[i], there contains character 'a' + i
	for i, w := range words {
		for _, char := range w {
			intDict[i] |= 1 << (char - 'a')
		}
	}

	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if intDict[i]&intDict[j] == 0 { // no duplicate position with 1s
				size := len(words[i]) * len(words[j])
				if size > max {
					max = size
				}
			}
		}
	}
	return max
}
