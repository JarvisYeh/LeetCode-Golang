package main

import "fmt"

// reverse each word in a string
func reverseWords557(s string) string {
	arr := []byte(s)
	i := 0
	for i < len(arr) {
		// move j to the first space after i
		j := i
		for j < len(arr) && arr[j] != ' ' {
			j++
		}

		// reverse the word
		reverse557(arr, i, j-1)

		// update i to the start of the next word
		i = j + 1
	}
	return string(arr)
}

func reverse557(arr []byte, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func main() {
	input := "I love Yahoo"
	fmt.Println(input)
	fmt.Println(reverseWords557(input))
}
