package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	s = s[1 : len(s)-1]
	fmt.Println(s)
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		E[i] = []int{u, v}
	}
	return solve(n, E)
}

func solve(n int, E [][]int) []int {
	// ans[1] = 1 (它时一个完全图）
	// ans[2] = 如果只有3个节点，且长度为3，那么就分成了2部分
	// 找到最长的路径的两个端点(x, y)
	// 如果k > dist(x, y), 那么所有的点都断开了
	// 需要知道，每一个点离他最远的距离，然后当k > dist 它就孤立了
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dist := make([][]Pair, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		dist[u] = make([]Pair, 2)
		dist[u][0] = Pair{0, u}
		dist[u][1] = Pair{0, u}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				if dist[v][0].first+1 >= dist[u][0].first {
					dist[u][1] = dist[u][0]
					dist[u][0] = Pair{dist[v][0].first + 1, v}
				} else if dist[v][0].first+1 >= dist[u][1].first {
					dist[u][1] = Pair{dist[v][0].first + 1, v}
				}
			}
		}
	}

	dfs(0, 0)

	var dfs2 func(p int, u int, ln int)

	dfs2 = func(p int, u int, ln int) {
		if p >= 0 {
			if ln+1 >= dist[u][0].first {
				dist[u][1] = dist[u][0]
				dist[u][0] = Pair{ln + 1, p}
			} else if ln+1 >= dist[u][1].first {
				dist[u][1] = Pair{ln + 1, p}
			}
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if v == dist[u][0].second {
					dfs2(u, v, dist[u][1].first)
				} else {
					dfs2(u, v, dist[u][0].first)
				}
			}
		}
	}

	dfs2(-1, 0, 0)

	ans := make([]int, n+1)
	for i := range n {
		d := dist[i][0].first
		ans[d+1]++
	}
	ans[1] = 1
	for i := 2; i <= n; i++ {
		ans[i] += ans[i-1]
		ans[i] = min(ans[i], n)
	}

	return ans[1:]
}

type Pair struct {
	first  int
	second int
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
