package main

// method 1
// use array to store the predix sum
// every node looks upwards to see how many sub-path (path including curr node) has sum = target
// TC: O(n*h)
// SC: O(h)
func pathSum437I(root *TreeNode, targetSum int) int {
	count := new(int)
	*count = 0
	helper437I(root, 0, []int{}, count, targetSum)
	return *count
}

func helper437I(root *TreeNode, currSum int, prefixSums []int, count *int, target int) {
	if root == nil {
		return
	}

	// curr level check upwards until root node
	currSum += root.Val
	if currSum == target {
		*count++
	}
	for i := 0; i < len(prefixSums); i++ {
		if currSum-prefixSums[i] == target {
			*count++
		}
	}

	helper437I(root.Left, currSum, append(prefixSums, currSum), count, target)
	helper437I(root.Right, currSum, append(prefixSums, currSum), count, target)
}

// method 2
// the only difference is use map to stores prefix sum rather than array
// so look up on all prefix takes only O(1) than O(h)
// no need to loop the array anymore
// use map stores the prefix sum frequency from root to curr downwards
// 	[root, i-th node] sum, i from [root, curr)
// 	key = sum, value = count of the sum
// TC: O(n*1)
// SC: O(h)
func pathSum437II(root *TreeNode, targetSum int) int {
	count := new(int)
	*count = 0
	// countMap stores
	// key = root to curr sum
	// value = count of same sum
	countMap := make(map[int]int)
	countMap[0] = 1
	helper437II(root, 0, countMap, count, targetSum)
	return *count
}

func helper437II(root *TreeNode, currSum int, countMap map[int]int, count *int, target int) {
	if root == nil {
		return
	}

	// current level check how many sub-path ended at current node
	// add to count
	currSum += root.Val
	if val, ok := countMap[currSum-target]; ok {
		*count += val
	}

	// update countMap with this level currSum and traverse left and right child
	val, ok := countMap[currSum]
	if ok {
		countMap[currSum] = val + 1
	} else {
		countMap[currSum] = 1
	}
	helper437II(root.Left, currSum, countMap, count, target)
	helper437II(root.Right, currSum, countMap, count, target)

	// change back countMap and return to previous level
	if ok {
		countMap[currSum]--
	} else {
		delete(countMap, currSum)
	}
}

func main() {
	root := FormTreeByLayer([]string{"1", "-2", "-3", "1", "3", "-2", "null", "-1"})
	pathSum437II(root, 3)
}
