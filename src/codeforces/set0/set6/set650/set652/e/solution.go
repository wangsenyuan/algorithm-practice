package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		var u, v, x int
		fmt.Fscan(reader, &u, &v, &x)
		edges[i] = []int{u, v, x}
	}
	var a, b int
	fmt.Fscan(reader, &a, &b)
	return solve(n, edges, a, b)
}

func solve(n int, edges [][]int, a int, b int) bool {
	m := len(edges)
	g := NewGraph(n, 2*m)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vis := make([]bool, n)
	stack := make([]int, n)
	var top int
	dfn := make([]int, n)
	low := make([]int, n)
	var timer int

	var id int

	belong := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		timer++
		dfn[u] = timer
		low[u] = timer
		stack[top] = u
		top++
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if dfn[v] == 0 {
				dfs(u, v)
				low[u] = min(low[u], low[v])
			} else if vis[v] {
				low[u] = min(low[u], dfn[v])
			}
		}

		if low[u] == dfn[u] {
			for top > 0 {
				v := stack[top-1]
				top--
				vis[v] = false
				belong[v] = id
				if u == v {
					break
				}
			}
			id++
		}
	}

	dfs(-1, a-1)

	val := make([]int, id)

	tr := NewGraph(id, 2*id)

	for _, e := range edges {
		u, v, x := e[0]-1, e[1]-1, e[2]
		// belong[u]
		if belong[u] == belong[v] {
			// 这个SCC中有遗迹
			val[belong[u]] |= x
		} else {
			// 这个x咋搞呢？
			if belong[u] < belong[v] {
				u, v = v, u
			}
			// belong[u] > belong[v]
			// 越靠近a, 它的id越大
			tr.AddEdge(belong[u], belong[v])
			// 只有通过了这个桥，才能获得
			val[belong[v]] |= x
		}
	}

	fa := make([]int, id)

	var dfs1 func(u int)

	dfs1 = func(u int) {
		for i := tr.nodes[u]; i > 0; i = tr.next[i] {
			v := tr.to[i]
			fa[v] = u
			dfs1(v)
		}
	}

	fa[belong[a-1]] = -1
	dfs1(belong[a-1])

	u := belong[b-1]

	for u != -1 {
		if val[u] == 1 {
			return true
		}
		u = fa[u]
	}

	return false
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
