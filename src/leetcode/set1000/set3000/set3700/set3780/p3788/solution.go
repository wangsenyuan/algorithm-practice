package p3788


func maximumScore(nums []int) int64 {
	n := len(nums)
	var sum int
	for _, v := range nums {
		sum += v
	}
	suf := make([]int, n + 1)
	suf[n] = 1 << 60
	best := -(1 << 60)
	for i := n - 1; i > 0; i-- {
		suf[i] = min(suf[i+1], nums[i])
		sum -= nums[i]
		best = max(best, sum - suf[i])
	}

	return int64(best)
}