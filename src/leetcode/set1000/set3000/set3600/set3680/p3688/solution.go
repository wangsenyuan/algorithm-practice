package p3690

import "slices"

func maxTotalValue(nums []int, k int) int64 {
	// 最大值和最小值组成的区间，肯定需要被选中
	// k -= x
	x := slices.Min(nums)
	y := slices.Max(nums)
	return int64(y-x) * int64(k)
}
