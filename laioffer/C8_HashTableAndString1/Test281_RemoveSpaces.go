package main

import "fmt"

// keep one space between words
// trim other spaces and all the head and tail spaces
func removeSpaces281(input string) string {
	arr := []rune(input)
	i, j := 0, 0
	for j < len(arr) {
		if arr[j] != ' ' {
			arr[i] = arr[j]
			i++
			j++
			// else arr[j] == ' '
		} else if i == 0 || arr[i-1] == ' ' {
			j++
		} else { // if i != 0 && arr[i - 1] != ' ', preserver the space
			arr[i] = arr[j]
			i++
			j++
		}
	}

	// if there are elements after trim, remove the last space
	if len(input) > 0 && arr[i-1] == ' ' {
		return string(arr[:i-1])
	} else {
		return string(arr[:i])
	}
}

// cut down all spaces to one space
func removeSpaces281II(input string) string {
	arr := []rune(input)
	i, j := 0, 0
	for j < len(arr) {
		if arr[j] != ' ' {
			arr[i] = arr[j]
			i++
			j++
			// else arr[j] == ' '
		} else if i != 0 && arr[i-1] == ' ' {
			j++
		} else { // if i == 0 || arr[i - 1] != ' ', preserver the space
			arr[i] = arr[j]
			i++
			j++
		}
	}

	return string(arr[:i])
}

func main() {
	fmt.Println(removeSpaces281II("start:"+"  I love you ") + ":end")
}
