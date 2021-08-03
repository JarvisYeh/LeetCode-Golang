package main

import "fmt"

func reverseWords186(s []byte) {
	// step 1: reverse the whole sentence
	reverse186(s, 0, len(s)-1)

	// step 2: reverse each words
	i := 0 // i always points to the start of each words
	for i < len(s) {
		j := i
		// j points to the first space after words
		for j < len(s) && s[j] != ' ' {
			j++
		}
		reverse186(s, i, j-1)
		i = j + 1
	}

}

func reverse186(s []byte, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	input := []byte("I love Yahoo")
	fmt.Println(string(input))
	reverseWords186(input)
	fmt.Println(string(input))
}
