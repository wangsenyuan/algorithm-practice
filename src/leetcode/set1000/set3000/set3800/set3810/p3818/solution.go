package p3818

func minimumPrefixLength(nums []int) int {
	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			return i + 1
		}
	}
	return 0
}
