package p3615

func maxLen(n int, edges [][]int, label string) int {
	g := make([][]bool, n)
	for i := range n {
		g[i] = make([]bool, n)
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u][v] = true
		g[v][u] = true
	}
	N := 1 << n
	dp := make([][][]int, N)
	for s := range N {
		dp[s] = make([][]int, n)
		for i := range n {
			dp[s][i] = make([]int, n)
			for j := range n {
				dp[s][i][j] = -2 * n
			}
		}
	}

	for i := range n {
		dp[1<<i][i][i] = 1
	}

	for i := range n {
		for j := i + 1; j < n; j++ {
			if i != j && g[i][j] && label[i] == label[j] {
				dp[1<<i|1<<j][i][j] = 2
			}
		}
	}

	for s := range N {
		for i := range n {
			if s&(1<<i) == 0 {
				continue
			}
			for j := i; j < n; j++ {
				if s&(1<<j) == 0 || dp[s][i][j] < 0 {
					continue
				}
				for u := range n {
					if s&(1<<u) > 0 || !g[u][i] {
						continue
					}
					for v := range n {
						if u == v || s&(1<<v) > 0 || !g[j][v] || label[u] != label[v] {
							continue
						}
						dp[s|1<<u|1<<v][min(u, v)][max(u, v)] = max(dp[s|1<<u|1<<v][min(u, v)][max(u, v)], dp[s][i][j]+2)
					}
				}
			}
		}
	}

	var ans int
	for s := range N {
		for i := range n {
			for j := i; j < n; j++ {
				ans = max(ans, dp[s][i][j])
			}
		}
	}
	return ans
}
