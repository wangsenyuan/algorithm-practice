package p3892

const inf = 1 << 60

func minOperations(nums []int, k int) int {
	if k == 0 {
		return 0
	}
	n := len(nums)

	if k > n/2 {
		return -1
	}

	play := func(incFirst bool) int {
		dp := make([][]int, k+1)
		ndp := make([][]int, k+1)
		for i := range k + 1 {
			dp[i] = make([]int, 2)
			ndp[i] = make([]int, 2)
			for j := range 2 {
				dp[i][j] = inf
				ndp[i][j] = inf
			}
		}
		dp[0][0] = 0
		if incFirst {
			dp[1][1] = max(0, max(nums[1], nums[n-1])+1-nums[0])
		}

		// dp[x][1] 表示到i为止，有x个peak，且第i个数是peak时
		for i := 1; i < n; i++ {
			for j := range k + 1 {
				nj := min(k, j+1)

				delta := max(0, max(nums[(i+1)%n], nums[i-1])+1-nums[i])
				if i < n-1 || !incFirst {
					ndp[nj][1] = min(ndp[nj][1], dp[j][0]+delta)
				}
				// 当前值不作为peak时
				ndp[j][0] = min(ndp[j][0], dp[j][0], dp[j][1])
			}
			for j := range k + 1 {
				for d := range 2 {
					dp[j][d] = ndp[j][d]
					ndp[j][d] = inf
				}
			}
		}

		if incFirst {
			return dp[k][0]
		}
		return min(dp[k][0], dp[k][1])
	}

	return min(play(true), play(false))
}
