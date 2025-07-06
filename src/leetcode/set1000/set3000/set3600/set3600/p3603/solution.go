package p3603

const inf = 1 << 60

func minCost(m int, n int, waitCost [][]int) int64 {
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = inf
		}
	}
	// dp[i][j] 表示离开时(i, j)时的cost
	// 只有在奇数秒的时候进入(i, j)，并在那里等一秒，然后又在奇数秒移动到下一个位置
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+waitCost[i][j]+(i+1)*(j+1))
			}
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+waitCost[i][j]+(i+1)*(j+1))
			}
		}
	}

	res := dp[m-1][n-1]
	res -= waitCost[m-1][n-1]

	return int64(res)
}
