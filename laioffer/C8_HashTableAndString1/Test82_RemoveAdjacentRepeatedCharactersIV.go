package main

import "fmt"

func deDup82(input string) string {
	arr := []rune(input)
	i, j := 0, 0
	for j < len(arr) {
		if i == 0 || arr[j] != arr[i-1] {
			arr[i] = arr[j]
			i++
			j++
		} else {
			for j < len(arr) && arr[j] == arr[i-1] {
				j++
			}
			i--
		}
	}
	return string(arr[:i])
}

func main() {
	fmt.Println(deDup82("abbbbbacc"))
}
