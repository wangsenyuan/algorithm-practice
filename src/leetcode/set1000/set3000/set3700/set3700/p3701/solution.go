package p3701

func longestSubsequence(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	if res != 0 {
		return len(nums)
	}
	// res = 0
	n := len(nums)
	var first int
	for first < n && nums[first] == 0 {
		first++
	}
	if first == n {
		return 0
	}
	ans := n - first - 1

	last := n - 1
	for nums[last] == 0 {
		last--
	}
	return max(ans, last)
}
