package p3584

func maximumProduct(nums []int, m int) int64 {
	if m == 1 {
		var res int
		for _, x := range nums {
			res = max(res, x*x)
		}
		return int64(res)
	}
	ans := -(1 << 60)
	best := []int{nums[0], nums[0]}
	n := len(nums)
	for i := m - 1; i < n; i++ {
		ans = max(ans, nums[i]*best[0], nums[i]*best[1])
		best[0] = max(best[0], nums[i-m+2])
		best[1] = min(best[1], nums[i-m+2])
	}
	return int64(ans)
}
