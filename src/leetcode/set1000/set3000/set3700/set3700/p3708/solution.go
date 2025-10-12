package p3708

func longestSubarray(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	best := 2
	cur := 2
	for i := 2; i < n; i++ {
		if nums[i] == nums[i-1]+nums[i-2] {
			cur++
		} else {
			cur = 2
		}
		best = max(best, cur)
	}

	return best
}
