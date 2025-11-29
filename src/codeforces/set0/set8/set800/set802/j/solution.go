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
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		edges[i] = []int{u, v, w}
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	g := NewGraph(n, 2*n)

	for _, cur := range edges {
		u, v, w := cur[0], cur[1], cur[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	var dfs func(p int, u int) int

	dfs = func(p int, u int) int {
		var res int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				res = max(res, dfs(u, v)+g.val[i])
			}
		}
		return res
	}

	return dfs(0, 0)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
