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
	var n, m, g int
	fmt.Fscan(reader, &n, &m, &g)
	sex := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &sex[i])
	}
	v := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &v[i])
	}
	folks := make([][]int, m)
	for i := range m {
		var e, w, k int
		fmt.Fscan(reader, &e, &w, &k)
		folks[i] = make([]int, 3+k+1)
		folks[i][0] = e
		folks[i][1] = w
		folks[i][2] = k
		for j := 3; j < len(folks[i]); j++ {
			fmt.Fscan(reader, &folks[i][j])
		}
	}

	return solve(n, m, g, sex, v, folks)
}

type RichMan struct {
	id     int
	w      int
	friend bool
	want   int
	dogs   []int
}

func solve(n int, m int, pay int, sex []int, v []int, folks [][]int) int {
	men := make([]RichMan, len(folks))
	var ans int
	var edge_count int
	for i, cur := range folks {
		ans += cur[1]
		k := len(cur)
		man := RichMan{id: i}
		man.w = cur[1]
		man.want = cur[0]
		man.friend = cur[k-1] == 1
		man.dogs = cur[3 : k-1]
		men[i] = man
		edge_count += k
	}

	g := NewGraph(n+m+2, 2*(n+edge_count))
	S := n + m
	T := S + 1

	addEdge := func(u, v, w int) {
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, 0)
	}

	for i := range n {
		if sex[i] == 0 {
			addEdge(S, i, v[i])
		} else {
			addEdge(i, T, v[i])
		}
	}

	for i := range m {
		w := men[i].w
		if men[i].friend {
			w += pay
		}
		if men[i].want == 0 {
			addEdge(S, n+i, w)
			for _, j := range men[i].dogs {
				j--
				addEdge(n+i, j, inf)
			}
		} else {
			for _, j := range men[i].dogs {
				j--
				addEdge(j, n+i, inf)
			}
			addEdge(n+i, T, w)
		}
	}

	mx := dinic(S, T, n+m+2, g)

	return ans - mx
}

const inf = 1 << 60

func dinic(src, snk int, n int, g *Graph) int {
	level := make([]int, n)
	all_minus_one := make([]int, n)
	for i := 0; i < n; i++ {
		all_minus_one[i] = -1
	}

	que := make([]int, n)

	bfs := func() bool {
		var front, end int
		copy(level, all_minus_one)
		level[src] = 0
		que[end] = src
		end++
		for front < end {
			u := que[front]
			front++
			for i := g.node[u]; i > 0; i = g.next[i] {
				if g.limit[i] > g.flow[i] && level[g.to[i]] == -1 {
					v := g.to[i]
					level[v] = level[u] + 1
					que[end] = v
					end++
				}
			}
		}
		return level[snk] > 0
	}

	pos := make([]int, n)

	var dfs func(u int, flow int) int
	dfs = func(u int, flow int) int {
		if flow == 0 {
			return 0
		}
		if u == snk {
			return flow
		}

		for pos[u] > 0 {
			i := pos[u]
			v := g.to[i]
			if level[v] == level[u]+1 && g.flow[i] < g.limit[i] {
				tr := dfs(v, min(flow, g.limit[i]-g.flow[i]))
				if tr > 0 {
					g.flow[i] += tr
					g.flow[i^1] -= tr
					return tr
				}
			}

			pos[u] = g.next[i]
		}
		return 0
	}
	var flow int
	for bfs() {
		for i := 0; i < n; i++ {
			pos[i] = g.node[i]
		}
		for {
			cur := dfs(src, inf)
			if cur == 0 {
				break
			}
			flow += cur
		}
	}
	return flow
}

type Graph struct {
	node  []int
	next  []int
	to    []int
	flow  []int
	limit []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.flow = make([]int, e+3)
	g.limit = make([]int, e+3)
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.limit[g.cur] = w
	g.flow[g.cur] = 0
}
