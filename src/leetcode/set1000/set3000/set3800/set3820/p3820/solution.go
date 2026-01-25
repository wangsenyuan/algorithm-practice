package p3820

func specialNodes(n int, edges [][]int, x int, y int, z int) int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	que := make([]int, n)
	bfs := func(s int) []int {
		dist := make([]int, n)
		for i := range n {
			dist[i] = -1
		}
		var head, tail int
		que[head] = s
		head++
		dist[s] = 0
		for tail < head {
			u := que[tail]
			tail++
			for _, v := range adj[u] {
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}
		return dist
	}

	dx := bfs(x)
	dy := bfs(y)
	dz := bfs(z)

	var res int
	for i := range n {
		a, b, c := dx[i], dy[i], dz[i]
		if a < b {
			a, b = b, a
		}
		// a > b
		if a < c {
			a, c = c, a
		}
		// a >= c
		if b*b+c*c == a*a {
			res++
		}
	}

	return res
}
