package p3742

const inf = 1 << 30

func maxPathScore(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	k = min(k, m+n-1)

	dp := make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, m)
		for j := range m {
			dp[i][j] = make([]int, k+1)
			for l := range k + 1 {
				dp[i][j][l] = -inf
			}
		}
	}
	dp[0][0][0] = 0

	for i := range n {
		for j := range m {
			var x int
			if grid[i][j] > 0 {
				x = 1
			}
			for l := x; l <= k; l++ {
				if i > 0 {
					dp[i][j][l] = max(dp[i][j][l], dp[i-1][j][l-x]+grid[i][j])
				}
				if j > 0 {
					dp[i][j][l] = max(dp[i][j][l], dp[i][j-1][l-x]+grid[i][j])
				}
				if l > 0 {
					dp[i][j][l] = max(dp[i][j][l], dp[i][j][l-1])
				}
			}
		}
	}
	if dp[n-1][m-1][k] < 0 {
		return -1
	}
	return dp[n-1][m-1][k]
}
