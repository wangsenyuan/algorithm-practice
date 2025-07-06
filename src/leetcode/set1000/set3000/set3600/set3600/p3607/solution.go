package p3607

import "sort"

func processQueries(c int, connections [][]int, queries [][]int) []int {
	n := len(connections)
	g := NewGraph(c+1, 2*n)
	for _, e := range connections {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vis := make([]bool, c+1)

	var arr []int
	var comps [][]int
	belonge := make([]int, c+1)

	var dfs func(u int)
	dfs = func(u int) {
		belonge[u] = len(comps)
		arr = append(arr, u)
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
	}

	for u := 1; u <= c; u++ {
		if !vis[u] {
			arr = make([]int, 0, 1)
			dfs(u)
			sort.Ints(arr)
			comps = append(comps, arr)
		}
	}

	var ans []int

	marked := make([]bool, c+1)

	pos := make([]int, len(comps))

	for _, cur := range queries {
		if cur[0] == 1 {
			x := cur[1]
			if !marked[x] {
				ans = append(ans, x)
			} else {
				j := belonge[x]
				for pos[j] < len(comps[j]) && marked[comps[j][pos[j]]] {
					pos[j]++
				}
				if pos[j] == len(comps[j]) {
					ans = append(ans, -1)
				} else {
					ans = append(ans, comps[j][pos[j]])
				}
			}
		} else {
			x := cur[1]
			marked[x] = true
		}
	}

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
	e++
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
