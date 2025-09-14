package p3684

import (
	"slices"
	"sort"
)

func maxKDistinct(nums []int, k int) []int {
	sort.Ints(nums)
	nums = slices.Compact(nums)
	slices.Reverse(nums)
	k = min(k, len(nums))
	return nums[:k]
}
