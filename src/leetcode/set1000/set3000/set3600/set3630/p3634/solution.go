package p3634

import "sort"

func minRemoval(nums []int, k int) int {
	sort.Ints(nums)

	n := len(nums)

	ans := n - 1

	for l, r := 0, 0; r < n; r++ {
		for l < r && nums[l]*k < nums[r] {
			l++
		}
		ans = min(ans, n-(r-l+1))
	}

	return ans
}
