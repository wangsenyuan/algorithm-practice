package p3965

func finishTime(n int, edges [][]int, baseTime []int) int64 {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	dp := make([]int, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		earliest, latest := 1<<60, 0
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v)
				earliest = min(earliest, dp[v])
				latest = max(latest, dp[v])
			}
		}
		if earliest <= latest {
			dp[u] = latest - earliest + baseTime[u] + latest
		} else {
			dp[u] = baseTime[u]
		}
	}

	dfs(-1, 0)

	return int64(dp[0])
}
