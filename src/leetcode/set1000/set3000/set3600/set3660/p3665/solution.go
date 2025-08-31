package p3665

const MOD = 1000000007

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func uniquePaths(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	// fp[i][j][0] 表示从右边进入时，能够到达的位置
	fp := make([][][2]int, n)
	dp := make([][]int, n)
	for i := range n {
		fp[i] = make([][2]int, m)
		dp[i] = make([]int, m)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			fp[i][j][0] = i*m + j
			fp[i][j][1] = i*m + j
			if grid[i][j] == 1 {
				if j+1 == m {
					// 如果从上进入(i, j)，会出界
					fp[i][j][1] = n * m
				} else {
					fp[i][j][1] = fp[i][j+1][0]
				}
				if i+1 == n {
					fp[i][j][0] = n * m
				} else {
					fp[i][j][0] = fp[i+1][j][1]
				}
			}
		}
	}

	dp[0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				if j+1 < m && fp[i][j+1][0] < n*m {
					r, c := fp[i][j+1][0]/m, fp[i][j+1][0]%m
					dp[r][c] = add(dp[r][c], dp[i][j])
				}
				if i+1 < n && fp[i+1][j][1] < n*m {
					r, c := fp[i+1][j][1]/m, fp[i+1][j][1]%m
					dp[r][c] = add(dp[r][c], dp[i][j])
				}
			}
		}
	}

	return dp[n-1][m-1]
}
