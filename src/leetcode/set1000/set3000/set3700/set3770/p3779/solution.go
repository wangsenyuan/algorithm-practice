package p3779

import "slices"

func minOperations(nums []int) int {
	x := slices.Max(nums)
	freq := make([]int, x+1)
	n := len(nums)

	for i := n - 1; i >= 0; i-- {
		freq[nums[i]]++
		if freq[nums[i]] > 1 {
			return (i + 3) / 3
		}
	}
	return 0
}
