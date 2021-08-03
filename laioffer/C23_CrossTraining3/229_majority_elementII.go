package main

func majorityElement229(nums []int) []int {
	res := []int{}
	can1, can2, count1, count2 := 0, 0, 0, 0
	for _, num := range nums {
		if num == can1 {
			count1++
		} else if num == can2 {
			count2++
		} else if count1 == 0 {
			can1 = num
			count1 = 1
		} else if count2 == 0 {
			can2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	// post-processing, count each candidates
	count1, count2 = 0, 0
	for _, num := range nums {
		if can1 == num {
			count1++
		}
		if can2 == num {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		res = append(res, can1)
	}
	if can2 != can1 && count2 > len(nums)/3 {
		res = append(res, can2)
	}
	return res
}

func main() {
	majorityElement229([]int{2, 1, 1, 3, 1, 4, 5, 6})
}
