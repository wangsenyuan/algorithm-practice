package p3772

func maxSubgraphScore(n int, edges [][]int, good []int) []int {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dp := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				dp[u] += max(0, dp[v])
			}
		}
		if good[u] == 1 {
			dp[u]++
		} else {
			dp[u]--
		}
	}

	dfs(-1, 0)

	ans := make([]int, n)

	var dfs2 func(p int, u int, x int)
	dfs2 = func(p int, u int, x int) {
		if p >= 0 {
			dp[u] += max(0, x)
		}
		ans[u] = dp[u]
		sum := dp[u]

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				w := dp[v]
				sum -= max(0, w)

				dfs2(u, v, sum)

				sum += max(0, w)
			}
		}
	}

	dfs2(-1, 0, 0)

	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
