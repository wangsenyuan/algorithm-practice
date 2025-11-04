package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	if len(res) == 0 {
		fmt.Fprintln(writer, -1)
		return
	}
	for _, row := range res {
		fmt.Fprintln(writer, len(row))
		for _, v := range row {
			fmt.Fprint(writer, v, " ")
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) [][]int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	return solve(n, edges)
}
func solve(n int, edges [][]int) [][]int {
	m := len(edges)
	g := NewGraph(n, 2*m)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	color := make([]int, n)
	for i := range n {
		color[i] = -1
	}

	var dfs func(u int, c int) bool
	dfs = func(u int, c int) bool {
		if color[u] != -1 {
			return color[u] == c
		}
		color[u] = c
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, 1^c) {
				return false
			}
		}
		return true
	}

	var set int

	for u := range n {
		if color[u] == -1 {
			if !dfs(u, set) {
				return nil
			}
			// 保证两个都是非空的
			set ^= 1
		}
	}

	res := make([][]int, 2)
	for i := range n {
		res[color[i]] = append(res[color[i]], i+1)
	}
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
