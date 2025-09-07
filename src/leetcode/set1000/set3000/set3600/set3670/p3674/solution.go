package p3674

import "slices"

func minOperations(nums []int) int {
	x := slices.Min(nums)
	y := slices.Max(nums)
	if x == y {
		return 0
	}
	return 1
}
