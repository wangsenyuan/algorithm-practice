package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) int {
	var n, m, s int
	fmt.Fscan(reader, &n, &m, &s)
	roads := make([][]int, m)
	for i := range m {
		roads[i] = readNNums(reader, 2)
	}
	return solve(n, m, s, roads)
}

func solve(n int, m int, s int, roads [][]int) int {
	// 先搞出scc
	g := NewGraph(n, m)
	for _, road := range roads {
		g.AddEdge(road[0]-1, road[1]-1)
	}

	dfn := make([]int, n)
	low := make([]int, n)
	var timer int
	stack := make([]int, n)
	var top int
	vis := make([]bool, n)

	id := make([]int, n)
	var scc int

	var dfs func(u int)
	dfs = func(u int) {
		timer++
		dfn[u] = timer
		low[u] = timer
		stack[top] = u
		top++
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dfn[v] == 0 {
				dfs(v)
				low[u] = min(low[u], low[v])
			} else if vis[v] {
				low[u] = min(low[u], dfn[v])
			}
		}
		if dfn[u] == low[u] {
			for top > 0 {
				v := stack[top-1]
				top--
				vis[v] = false
				id[v] = scc
				if u == v {
					break
				}
			}
			scc++
		}
	}

	for u := range n {
		if dfn[u] == 0 {
			dfs(u)
		}
	}

	deg := make([]int, scc)
	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		if id[u] != id[v] {
			deg[id[v]]++
		}
	}

	// 现在形成了一个dag, deg[?] = 0的节点，都需要有一条边连接到
	var ans int
	for i := range scc {
		if deg[i] == 0 {
			ans++
		}
	}

	if deg[id[s-1]] == 0 {
		ans--
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
