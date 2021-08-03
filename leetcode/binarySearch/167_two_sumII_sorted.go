package main

// use binary search
// TC: O(log n)
// SC: O(1)
func twoSum167I(numbers []int, target int) []int {
	for i, _ := range numbers {
		t := target - numbers[i]
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := low + ((high - low) >> 1)
			if t == numbers[mid] {
				return []int{i + 1, mid + 1}
			}
			if numbers[mid] > t {
				high = mid - 1
			} else if numbers[mid] < t {
				low = mid + 1
			} else {
				break
			}
		}

	}
	return []int{-1, -1}
}

// use two pointer
// TC: O(n)
// SC: O(1)
func twoSum167II(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{-1, -1}
}

// use map
// TC: O(n)
// SC: O(n)
func twoSum167III(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		another := target - numbers[i]
		if val, found := m[another]; found {
			return []int{val + 1, i + 1}
		} else {
			m[numbers[i]] = i
		}
	}
	return []int{-1, -1}
}

func main() {
	twoSum167III([]int{0, 0, 3, 4}, 0)
}
