package p3724

func minOperations(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	var res int
	best := 1 << 30
	for i := range n {
		res += abs(nums1[i] - nums2[i])
		best = min(best, abs(nums2[i]-nums2[n]))
		best = min(best, abs(nums1[i]-nums2[n]))
		if nums1[i] <= nums2[n] && nums2[n] <= nums2[i] {
			best = 0
		}
		if nums1[i] >= nums2[n] && nums2[n] >= nums2[i] {
			best = 0
		}
	}

	res += best + 1

	return int64(res)
}

func abs(num int) int {
	return max(num, -num)
}
