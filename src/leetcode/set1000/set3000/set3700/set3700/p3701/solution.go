package p3701

import "slices"

func longestSubsequence(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	if res != 0 {
		return len(nums)
	}
	// res = 0
	n := len(nums)
	x := slices.Max(nums)
	if x == 0 {
		return 0
	}
	return n - 1
}
