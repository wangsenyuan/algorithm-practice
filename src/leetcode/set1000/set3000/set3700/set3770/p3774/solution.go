package p3774

import "slices"

func absDifference(nums []int, k int) int {
	slices.Sort(nums)
	var sum1 int
	for i := range k {
		sum1 += nums[i]
	}
	n := len(nums)
	var sum2 int
	for i := n - k; i < n; i++ {
		sum2 += nums[i]
	}
	return sum2 - sum1
}
