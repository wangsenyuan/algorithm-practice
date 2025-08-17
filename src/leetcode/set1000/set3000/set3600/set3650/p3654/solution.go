package p3654

const inf = 1 << 60

func minArraySum(nums []int, k int) int64 {
	// s1 % k = 0, s2 % k
	dp := make([]int, k)
	for i := range k {
		dp[i] = inf
	}
	dp[0] = 0
	var sum int
	pos := make([]int, k)
	for i := range k {
		pos[i] = -1
	}
	pos[0] = 0
	n := len(nums)

	fp := make([]int, n+1)
	for i, v := range nums {
		i++
		// 这个是不删除i时的最小值
		fp[i] = dp[sum] + v
		sum = (sum + v) % k
		if pos[sum] >= 0 {
			// 把这段删除掉
			fp[i] = min(fp[i], fp[pos[sum]])
		}
		dp[sum] = min(dp[sum], fp[i])
		pos[sum] = i
	}

	return int64(dp[sum])
}
