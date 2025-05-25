package p3562

func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	g := NewGraph(n, n)

	for _, cur := range hierarchy {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v)
	}

	var dfs func(u int) [][2]int

	dfs = func(u int) [][2]int {
		sub := make([][2]int, budget+1)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			tmp := dfs(v)
			for j := budget; j >= 0; j-- {
				for jv := 0; jv <= j; jv++ {
					for k := range 2 {
						sub[j][k] = max(sub[j][k], sub[j-jv][k]+tmp[jv][k])
					}
				}
			}
		}
		f := make([][2]int, budget+1)

		for j := 0; j <= budget; j++ {
			for k := range 2 {
				cost := present[u] / (k + 1)
				if j >= cost {
					f[j][k] = max(sub[j][0], sub[j-cost][1]+future[u]-cost)
				} else {
					f[j][k] = sub[j][0]
				}
			}
		}

		return f
	}

	res := dfs(0)
	return res[budget][0]
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
