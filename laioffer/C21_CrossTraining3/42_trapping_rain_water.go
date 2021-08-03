package main

// TC: O(n)
// SC: O(n)
func trap42(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	n := len(height)

	// leftMax[i]: the largest height look left from i-th index
	// rightMax[i]: the largest height look right from i-th index
	leftMax, rightMax := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 || height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}

	for i := n - 1; i >= 0; i-- {
		if i == n-1 || height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	max := 0
	for i := 0; i < n; i++ {
		if leftMax[i] > rightMax[i] {
			max += (rightMax[i] - height[i])
		} else {
			max += (leftMax[i] - height[i])
		}
	}
	return max
}

// TC: O(n)
// SC: O(1)
func trap(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	n := len(height)
	leftMax, rightMax := height[0], height[n-1]
	i, j := 0, n-1
	// based on the fact that
	// for i-th element, leftMax is certain, rightMax might be larger
	// for j-th element, leftMax might be larger, rightMax is ceratin
	// and for a elemnet, the min(leftMax, rightMax) determines the actual height
	sum := 0
	for i <= j {
		if leftMax >= rightMax {
			// for j-th element, since rightMax is certain, and leftMax will only be larger
			// height is determined by rightMax
			if height[j] < rightMax {
				sum += rightMax - height[j]
			} else {
				rightMax = height[j]
			}
			j--
		} else {
			// for i-th element, since leftMax is ceratin, and rightMax will only be larger
			// height is determined by leftMax
			if height[i] < leftMax {
				sum += leftMax - height[i]
			} else {
				leftMax = height[i]
			}
			i++
		}
	}
	return sum
}
