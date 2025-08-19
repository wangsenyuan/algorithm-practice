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
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	if n%2 == 1 {
		return -1
	}
	g := NewGraph(n, 2*n)
	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}

	sz := make([]int, n)
	marked := make([]bool, n)
	que := make([]int, n)
	var head, tail int
	for u := range n {
		sz[u] = 1
		if deg[u] == 1 {
			marked[u] = true
			que[head] = u
			head++
		}
	}
	var res int

	var rem int

	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !marked[v] {
				sz[v] += sz[u]
				deg[v]--
				if deg[v] == 1 {
					marked[v] = true
					que[head] = v
					head++
					if sz[v]%2 == 0 {
						res++
						rem += sz[v]
						sz[v] = 0
					}
				}
			}
		}
	}

	if (n-rem)%2 != 0 {
		return -1
	}

	if rem == n {
		res--
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
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
