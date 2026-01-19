package p3810

import "slices"

func minOperations(nums []int, target []int) int {
	n := len(nums)
	x := slices.Max(nums)

	marked := make([]bool, x+1)
	for i := range n {
		if nums[i] != target[i] {
			marked[nums[i]] = true
		}
	}

	var res int
	for i := range x + 1 {
		if marked[i] {
			res++
		}
	}

	return res
}
