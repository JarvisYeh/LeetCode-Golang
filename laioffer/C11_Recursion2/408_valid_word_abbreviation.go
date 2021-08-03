package main

func validWordAbbreviation408I(word string, abbr string) bool {
	i, j := 0, 0
	for i < len(word) && j < len(abbr) {
		if abbr[j] >= 'a' && abbr[j] <= 'z' { // it is a character
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++
		} else if abbr[j] >= '0' && abbr[j] <= '9' { // it is a digit
			count := 0
			// handle special case ("a", "01") should return false in test case
			if abbr[j] == '0' {
				return false
			}
			for j < len(abbr) && abbr[j] >= '0' && abbr[j] <= '9' {
				count *= 10
				count += int(abbr[j] - '0')
				j++
			}
			i += count
		}

	}

	if i == len(word) && j == len(abbr) {
		return true
	}

	return false
}

func validWordAbbreviation408II(word string, abbr string) bool {
	return checkValid408II(word, abbr, 0, 0)
}

func checkValid408II(word, abbr string, i, j int) bool {
	// reach end at same time
	if i == len(word) && j == len(abbr) {
		return true
	}

	// reach end at different time
	if i == len(word) || j == len(abbr) {
		return false
	}

	// i overflows before j reaches end
	if i > len(word) {
		return false
	}

	// if is alphabetic
	if abbr[j] >= 'a' && abbr[j] <= 'z' {
		if abbr[j] == word[i] {
			return checkValid408II(word, abbr, i+1, j+1)
		} else {
			return false
		}
	}

	// if is digit
	if abbr[j] >= '0' && abbr[j] <= '9' {
		if abbr[j] == '0' {
			return false
		}
		count := 0
		for j < len(abbr) && abbr[j] >= '0' && abbr[j] <= '9' {
			count *= 10
			count += int(abbr[j] - '0')
			j++
		}
		return checkValid408II(word, abbr, i+count, j)
	}
	return false
}

func main() {
	validWordAbbreviation408II("internationalization", "i12iz4n")
}
