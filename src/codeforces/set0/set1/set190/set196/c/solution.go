package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func process(reader *bufio.Reader) (n int, edges [][]int, points [][]int, res []int) {
	fmt.Fscan(reader, &n)
	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	points = make([][]int, n)
	for i := range n {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	res = solve(n, edges, points)
	return
}

type point struct {
	id int
	x  int
	y  int
}

func cross(a, b, c point) int {
	return (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x)
}

func solve(n int, edges [][]int, points [][]int) []int {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	ans := make([]int, n)

	sz := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0)

	var assign func(p int, u int, arr []int)

	assign = func(p int, u int, arr []int) {
		o := 0
		for i := 1; i < len(arr); i++ {
			p1 := points[arr[o]]
			p2 := points[arr[i]]
			if p2[0] < p1[0] || p2[0] == p1[0] && p2[1] < p1[1] {
				o = i
			}
		}

		ans[arr[o]] = u + 1

		base := points[arr[o]]

		tmp := slices.Clone(arr)
		copy(tmp[o:], tmp[o+1:])

		tmp = tmp[:len(tmp)-1]

		slices.SortFunc(tmp, func(i, j int) int {
			a := points[i]
			b := points[j]
			// Cross product to determine angle relative to base
			c := (a[0]-base[0])*(b[1]-base[1]) - (a[1]-base[1])*(b[0]-base[0])
			return -c
		})

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				assign(u, v, tmp[:sz[v]])
				tmp = tmp[sz[v]:]
			}
		}
	}

	all := make([]int, n)
	for i := range n {
		all[i] = i
	}

	assign(0, 0, all)

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
