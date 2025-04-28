package p3528

const mod = 1e9 + 7

func mul(a, b int) int {
	return a * b % mod
}

func baseUnitConversions(conversions [][]int) []int {
	// bfs or dfs
	n := len(conversions) + 1

	type edge struct {
		to     int
		weight int
	}

	g := make([][]edge, n)
	for _, cur := range conversions {
		u, v, w := cur[0], cur[1], cur[2]
		g[u] = append(g[u], edge{v, w})
	}
	ans := make([]int, n)

	ans[0] = 1

	que := make([]int, n)
	var head, tail int
	que[head] = 0
	head++
	for tail < head {
		u := que[tail]
		tail++
		for _, cur := range g[u] {
			v, w := cur.to, cur.weight
			ans[v] = mul(ans[u], w)
			que[head] = v
			head++
		}
	}

	return ans
}
