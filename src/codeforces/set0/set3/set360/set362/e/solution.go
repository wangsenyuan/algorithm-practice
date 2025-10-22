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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(n, k, a)
}

func solve(n int, k int, a [][]int) int {
	src := 0
	snk := n - 1

	// Binary search on the flow amount
	check := func(expect int) (int, int) {
		// Rebuild graph for each check (like C++ version)
		g := NewGraph(n, n*n*4)

		addEdge := func(u, v, limit, cost int) {
			g.AddEdge(u, v, limit, cost)
			g.AddEdge(v, u, 0, -cost)
		}

		for i := range n {
			for j := range n {
				if i != j && a[i][j] > 0 {
					addEdge(i, j, a[i][j], 0)
					addEdge(i, j, INF, 1)
				}
			}
		}

		dist := make([]int, n)
		inQueue := make([]bool, n)

		spfa := func() bool {
			for i := range n {
				dist[i] = INF
				inQueue[i] = false
			}
			dist[src] = 0
			queue := []int{src}
			inQueue[src] = true

			for len(queue) > 0 {
				u := queue[0]
				queue = queue[1:]
				inQueue[u] = false

				for i := g.node[u]; i > 0; i = g.next[i] {
					v := g.to[i]
					if g.limit[i] > g.flow[i] && dist[v] > dist[u]+g.val[i] {
						dist[v] = dist[u] + g.val[i]
						if !inQueue[v] {
							queue = append(queue, v)
							inQueue[v] = true
						}
					}
				}
			}
			return dist[snk] < INF
		}

		vis := make([]bool, n)
		cur := make([]int, n)
		cost := 0

		var dfs func(u int, flow int) int
		dfs = func(u int, flow int) int {
			if u == snk {
				vis[snk] = true
				return flow
			}
			vis[u] = true
			sum := 0
			for cur[u] > 0 {
				i := cur[u]
				v := g.to[i]
				if g.limit[i] > g.flow[i] && dist[v] == dist[u]+g.val[i] && (!vis[v] || v == snk) {
					pushed := dfs(v, min(flow, g.limit[i]-g.flow[i]))
					if pushed > 0 {
						cost += g.val[i] * pushed
						g.flow[i] += pushed
						g.flow[i^1] -= pushed
						flow -= pushed
						sum += pushed
						if flow == 0 {
							break
						}
					}
				}
				cur[u] = g.next[i]
			}
			// 这个有点奇怪的，貌似是记录了这个节点上，可以流出去的总量～
			if sum == 0 {
				dist[u] = -INF
			}
			return sum
		}

		totalFlow := 0

		for totalFlow < expect && spfa() {
			for i := range n {
				cur[i] = g.node[i]
				vis[i] = false
			}
			pushed := dfs(src, expect-totalFlow)
			if pushed == 0 {
				break
			}
			totalFlow += pushed
		}

		return totalFlow, cost
	}

	l, r := 0, INF
	ans := 0
	for l <= r {
		mid := (l + r) / 2
		flow, cost := check(mid)
		if flow == 0 {
			return 0
		}
		if cost <= k {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}

const INF = 1 << 30

type Graph struct {
	node  []int
	next  []int
	to    []int
	flow  []int
	limit []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.flow = make([]int, e+3)
	g.limit = make([]int, e+3)
	g.val = make([]int, e+3)
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v, c, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.limit[g.cur] = c
	g.flow[g.cur] = 0
	g.val[g.cur] = w
}
