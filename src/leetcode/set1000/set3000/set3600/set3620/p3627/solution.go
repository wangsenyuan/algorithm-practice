package p3627

import "sort"

func maximumMedianSum(nums []int) int64 {
	sort.Ints(nums)

	n := len(nums)
	var res int

	lo, hi := 0, n

	for lo+3 <= hi {
		res += nums[hi-2]
		hi -= 2
		lo++
	}
	return int64(res)
}
