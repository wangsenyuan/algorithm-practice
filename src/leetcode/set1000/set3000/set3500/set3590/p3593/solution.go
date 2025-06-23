package p3593

func minIncrease(n int, edges [][]int, cost []int) int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	val := make([]int, n)

	var res int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		var mx int
		var cnt int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				cnt++
				dfs(u, v)
				mx = max(mx, val[v])
			}
		}
		if cnt > 0 {
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p != v && val[v] == mx {
					cnt--
				}
			}
			res += cnt
		}
		val[u] = mx + cost[u]
	}

	dfs(0, 0)

	return res
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
