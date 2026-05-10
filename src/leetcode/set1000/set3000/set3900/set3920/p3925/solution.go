package p3925

import "slices"


func concatWithReverse(nums []int) []int {
	n := len(nums)
	res := make([]int, 2 * n)
	copy(res, nums)
	slices.Reverse(nums)
	copy(res[n:], nums)
	return res
}