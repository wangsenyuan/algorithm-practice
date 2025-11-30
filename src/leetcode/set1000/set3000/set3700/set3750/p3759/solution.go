package p3759

import "slices"

func countElements(nums []int, k int) int {
	n := len(nums)
	if k == 0 {
		return n
	}
	slices.Sort(nums)
	w := nums[n-k]
	// res := n - k
	for i := n - k - 1; i >= 0; i-- {
		if nums[i] < w {
			return i + 1
		}
	}
	return 0
}
