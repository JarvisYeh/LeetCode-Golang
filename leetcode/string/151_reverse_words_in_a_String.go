package main

import "fmt"

// reverse the word order of string
// and remove extra spaces
func reverseWords151(s string) string {
	arr := []byte(s)
	// step 1: reserve whole string
	reverse151(arr, 0, len(arr)-1)

	// step 2: reverse each word
	reverseEachWords151(arr)

	// step 3: remove extra spaces
	end := cleanSpaces151(arr)
	return string(arr[:end])
}

func reverse151(arr []byte, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func reverseEachWords151(arr []byte) {
	i := 0
	for i < len(arr) {
		// move i to the start of the word
		for i < len(arr) && arr[i] == ' ' {
			i++
		}

		// move j to the first space after the word
		j := i
		for j < len(arr) && arr[j] != ' ' {
			j++
		}

		// if the start of the word exist, reverse the word
		if i < len(arr) {
			reverse151(arr, i, j-1)
		}
		i = j
	}
}

func cleanSpaces151(arr []byte) int {
	i, j := 0, 0
	for j < len(arr) {
		if arr[j] != ' ' {
			arr[i] = arr[j]
			i++
			j++
		} else if i != 0 && arr[i-1] != ' ' {
			arr[i] = arr[j]
			i++
			j++
		} else { // i == 0 || arr[i - 1] == ' '
			j++
		}
	}
	if arr[i-1] == ' ' {
		return i - 1
	} else {
		return i
	}
}

func main() {
	input := "   I  love  Yahoo    "
	fmt.Println(input)
	fmt.Println(reverseWords151(input))
}
