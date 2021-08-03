package main

func removeDuplicates1209(s string, k int) string {
	i, j := 0, 0
	arr := []rune(s)
	for j < len(arr) {
		if i == 0 || arr[j] != arr[i-1] {
			arr[i] = arr[j]
			i++
			j++
		} else {
			left := countDuplicateLeft1209(arr, i-1)
			right := countDuplicateRight1209(arr, j)
			if left+right < k { // reserve all duplicates
				for j < len(arr) && arr[j] == arr[i-1] {
					arr[i] = arr[j]
					j++
					i++
				}
			} else { // remove k duplicates
				i -= left
				j += (k - left)
			}
		}
	}
	return string(arr[:i])
}

func countDuplicateLeft1209(arr []rune, i int) int {
	tmp := arr[i]
	count := 0
	for ; i >= 0 && arr[i] == tmp; i-- {
		count++
	}
	return count
}

func countDuplicateRight1209(arr []rune, i int) int {
	tmp := arr[i]
	count := 0
	for ; i < len(arr) && arr[i] == tmp; i++ {
		count++
	}
	return count
}

func main() {
	removeDuplicates1209("pbbcggttciiippooaais", 2)
}
