package p3974

import "slices"

func maxSum(nums []int, k int, mul int) int64 {
	slices.Sort(nums)
	slices.Reverse(nums)
	// 选择k个数
	var res int64

	for i := range k {
		if mul > 0 {
			res += int64(nums[i] * mul)
		} else {
			res += int64(nums[i])
		}
		mul--
	}
	return res
}
