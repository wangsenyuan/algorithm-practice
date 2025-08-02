package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n, m, d int
	fmt.Fscan(reader, &n, &m, &d)
	p := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &p[i])
	}
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(d, n, edges, p)
}

type pair struct {
	first  int
	second int
}

func solve(d int, n int, edges [][]int, p []int) int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	marked := make([]bool, n)
	for _, v := range p {
		marked[v-1] = true
	}

	far := make([][2]pair, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		far[u] = [2]pair{{-1, -1}, {-1, -1}}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				if far[v][0].first >= 0 {
					if far[v][0].first+1 >= far[u][0].first {
						far[u][1] = far[u][0]
						far[u][0] = pair{far[v][0].first + 1, v}
					} else if far[v][0].first+1 >= far[u][1].first {
						far[u][1] = pair{far[v][0].first + 1, v}
					}
				}
			}
		}
		if marked[u] && far[u][0].first == -1 {
			far[u][0] = pair{0, u}
		}
	}

	dfs(-1, 0)

	var ans int

	var dfs2 func(p int, u int, f int)
	dfs2 = func(p int, u int, f int) {
		if p >= 0 && f >= 0 {
			if f+1 >= far[u][0].first {
				far[u][1] = far[u][0]
				far[u][0] = pair{f + 1, p}
			} else if f+1 >= far[u][1].first {
				far[u][1] = pair{f + 1, p}
			}
		}
		// far[u][0].first 不可能为-1, 因为肯定存在一个
		if far[u][0].first <= d {
			ans++
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				nf := far[u][0].first
				if far[u][0].second == v {
					nf = far[u][1].first
				}
				if nf < 0 && marked[u] {
					nf = 0
				}

				dfs2(u, v, nf)
			}
		}
	}

	dfs2(-1, 0, -1)

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
