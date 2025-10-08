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

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	roads := make([][]int, m)
	for i := 0; i < m; i++ {
		roads[i] = make([]int, 2)
		fmt.Fscan(reader, &roads[i][0], &roads[i][1])
	}
	first := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(reader, &first[i])
	}
	second := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(reader, &second[i])
	}
	return solve(n, roads, first, second)
}

func solve(n int, roads [][]int, first []int, second []int) int {
	g := NewGraph(n, len(roads)*2)
	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	que := make([]int, n)
	bfs := func(s int) []int {
		dist := make([]int, n)
		for i := range n {
			dist[i] = -1
		}
		dist[s] = 0
		var head, tail int
		que[head] = s
		head++
		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}
		return dist
	}

	dists := make([][]int, n)
	for i := range n {
		dists[i] = bfs(i)
	}
	s1, t1, l1 := first[0]-1, first[1]-1, first[2]
	s2, t2, l2 := second[0]-1, second[1]-1, second[2]
	if dists[s1][t1] > l1 || dists[s2][t2] > l2 {
		return -1
	}

	m := len(roads)
	best := m - dists[s1][t1] - dists[s2][t2]

	for u := range n {
		for v := range n {
			for i := range 2 {
				for j := range 2 {
					s1, t1 := first[i]-1, first[1-i]-1
					s2, t2 := second[j]-1, second[1-j]-1
					if dists[u][v]+dists[s1][u]+dists[v][t1] <= l1 && dists[u][v]+dists[s2][u]+dists[v][t2] <= l2 {
						best = max(best, m-dists[u][v]-dists[s1][u]-dists[s2][u]-dists[v][t1]-dists[v][t2])
					}
				}
			}

		}
	}

	return best

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
