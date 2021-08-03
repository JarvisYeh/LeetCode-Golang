package main

import "fmt"

func deDup79(input string) string {
	arr := []byte(input)
	i, j := 0, 0
	for j < len(arr) {
		if i == 0 || arr[j] != arr[i-1] {
			arr[i] = arr[j]
			i++
			j++
		} else {
			j++
		}
	}
	return string(arr[:i])
}

func main() {
	fmt.Println(deDup79("aabbccsdfff"))
}
