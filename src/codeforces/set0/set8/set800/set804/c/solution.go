package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, tot, assign := drive(reader)
	fmt.Println(tot)
	s := fmt.Sprintf("%v", assign)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (s [][]int, tot int, assign []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	s = make([][]int, n)
	for i := range n {
		var k int
		fmt.Fscan(reader, &k)
		s[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &s[i][j])
		}
	}
	edges := make([][]int, n-1)
	for i := range n - 1 {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	tot, assign = solve(n, m, s, edges)
	return
}

func solve(n int, m int, s [][]int, edges [][]int) (tot int, assign []int) {
	assign = make([]int, m)

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	occ := make([]bool, m+1)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for _, i := range s[u] {
			i--
			occ[assign[i]] = true
		}
		id := 1
		for _, i := range s[u] {
			i--
			if assign[i] == 0 {
				for occ[id] {
					id++
				}
				assign[i] = id
				occ[id] = true
			}
		}
		for _, i := range s[u] {
			i--
			occ[assign[i]] = false
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}
	}
	dfs(-1, 0)

	for i := range m {
		// 没有出现
		if assign[i] == 0 {
			assign[i] = 1
		}
	}

	return slices.Max(assign), assign
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
