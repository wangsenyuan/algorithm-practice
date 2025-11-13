package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	fmt.Println(res[0], res[1], res[2], res[3])
}

func drive(reader *bufio.Reader) (n int, roads [][]int, res []int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	roads = make([][]int, m)
	for i := range m {
		roads[i] = make([]int, 2)
		for j := range 2 {
			fmt.Fscan(reader, &roads[i][j])
		}
	}
	res = solve(n, roads)
	return
}

func solve(n int, roads [][]int) []int {
	m := len(roads)
	g := NewGraph(n, m)
	r := NewGraph(n, m)
	for _, cur := range roads {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v)
		r.AddEdge(v, u)
	}

	// 找到离s最远的三个city，以及距离
	que := make([]int, n)
	bfs := func(g *Graph, s int) (ans [][]int, dist []int) {
		dist = make([]int, n)
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
		ans = make([][]int, 3)
		for i := range 3 {
			ans[i] = make([]int, 2)
			ans[i][0] = s
		}
		var cnt int
		for i := range n {
			if dist[i] > 0 {
				cnt++
			}
			if dist[i] >= ans[0][1] {
				copy(ans[2], ans[1])
				copy(ans[1], ans[0])
				ans[0][0] = i
				ans[0][1] = dist[i]
			} else if dist[i] >= ans[1][1] {
				copy(ans[2], ans[1])
				ans[1][0] = i
				ans[1][1] = dist[i]
			} else if dist[i] >= ans[2][1] {
				ans[2][0] = i
				ans[2][1] = dist[i]
			}
		}
		if cnt < 3 {
			ans = ans[:cnt]
		}
		return
	}

	best1 := make([][][]int, n)
	d1 := make([][]int, n)
	best2 := make([][][]int, n)
	for i := range n {
		best1[i], d1[i] = bfs(g, i)
		best2[i], _ = bfs(r, i)
	}

	ans := make([]int, 5)

	for b := range n {
		for c := range n {
			if b == c || d1[b][c] == -1 || len(best2[b]) == 0 || len(best1[c]) == 0 {
				continue
			}
			// b 可以到达c

			if best2[b][0][1]+best1[c][0][1]+d1[b][c] < ans[4] {
				continue
			}

			for _, a := range best2[b] {
				for _, d := range best1[c] {
					cities := map[int]bool{a[0]: true, b: true, c: true, d[0]: true}
					if len(cities) == 4 {
						tmp := d1[b][c] + a[1] + d[1]
						if tmp > ans[4] {
							ans[0] = a[0] + 1
							ans[1] = b + 1
							ans[2] = c + 1
							ans[3] = d[0] + 1
							ans[4] = tmp
						}
					}
				}
			}
		}
	}

	return ans[:4]
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
