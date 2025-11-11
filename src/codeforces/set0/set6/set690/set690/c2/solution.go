package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	connectors := make([][]int, m)
	for i := range m {
		connectors[i] = make([]int, 2)
		fmt.Fscan(reader, &connectors[i][0], &connectors[i][1])
	}
	return solve(n, connectors)
}

func solve(n int, connectors [][]int) int {
	// a tree
	g := NewGraph(n, 2*n)
	for _, cur := range connectors {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var dfs func(p int, u int, d int) []int
	dfs = func(p int, u int, d int) []int {
		res := []int{d, u}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp := dfs(u, v, d+1)
				if tmp[0] > res[0] {
					res = tmp
				}
			}
		}
		return res
	}

	first := dfs(-1, 0, 0)
	second := dfs(-1, first[1], 0)

	return second[0]
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
