package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, _, res := drive(reader)
	fmt.Println(len(res))
	if len(res) > 0 {
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
}

func drive(reader *bufio.Reader) (n int, d int, edges [][]int, stations []int, res []int) {
	var k int
	fmt.Fscan(reader, &n, &k, &d)

	stations = make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &stations[i])
	}

	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	res = solve(d, n, edges, stations)
	return
}

func solve(d int, n int, edges [][]int, stations []int) []int {
	g := NewGraph(n, 2*n)

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}
	dist := make([]int, n)
	for i := range n {
		dist[i] = -1
	}

	que := make([]int, n)
	var head, tail int
	for _, x := range stations {
		x--
		if dist[x] == -1 {
			dist[x] = 0
			que[head] = x
			head++
		}
	}

	marked := make([]bool, len(edges))

	for tail < head {
		u := que[tail]
		tail++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				marked[g.val[i]] = true
				dist[v] = dist[u] + 1
				que[head] = v
				head++
			}
		}
	}

	var res []int
	for i := range n - 1 {
		if !marked[i] {
			res = append(res, i+1)
		}
	}
	return res
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
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
