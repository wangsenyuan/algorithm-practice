package p3727

import "slices"

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	for i := range n {
		if nums[i] < 0 {
			nums[i] = -nums[i]
		}
	}
	slices.Sort(nums)
	var res int
	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		res += nums[j] * nums[j]
		if i < j {
			res -= nums[i] * nums[i]
		}
	}
	return int64(res)
}
