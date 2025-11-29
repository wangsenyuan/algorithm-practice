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
	a := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}
	return solve(k, a, c)
}

const M = 1 << 30 // Huge number to ensure all (2i-1, 2i) edges are taken

func solve(k int, a []int, c []int) int {
	n := len(a)

	marked := make([]bool, n+1)
	var m int
	for i := range n {
		if !marked[a[i]] {
			m++
			marked[a[i]] = true
		}
	}

	k = min(k, m)

	var edges []Edge
	var src int
	dst := 2*n + 1

	for i := 1; i <= n; i++ {
		// From source to 2i-1 with cost c[a[i]]
		edges = append(edges, Edge{src, 2*i - 1, 1, c[a[i-1]-1]})
		// From 2i-1 to 2i with cost -M (huge negative to ensure this edge is always taken)
		edges = append(edges, Edge{2*i - 1, 2 * i, 1, -M})
		// From 2i to sink with cost 0
		edges = append(edges, Edge{2 * i, dst, 1, 0})

		// For each two requests i < j: edge from 2i to 2j-1
		for j := i + 1; j <= n; j++ {
			if a[i-1] == a[j-1] {
				// Same book: cost 0 (no need to buy again)
				edges = append(edges, Edge{2 * i, 2*j - 1, 1, 0})
			} else {
				// Different book: cost c[a[j]] (need to buy the book)
				edges = append(edges, Edge{2 * i, 2*j - 1, 1, c[a[j-1]-1]})
			}
		}
	}

	// Find minimum-cost flow of value exactly k
	// According to the solution: "any optimum solution must necessarily take all the edges (2i-1, 2i)"
	// because their cost is hugely negative. So we need k paths that together use all n edges.
	// The flow cost includes -n*M from the n edges (2i-1, 2i), so we add n*M to get actual cost.
	cost := minCostFlow(dst+1, edges, k, src, dst)
	// Add back n*M to compensate for the -M cost on each of the n edges (2i-1, 2i) that are taken
	return cost + n*M
}

type Edge struct {
	from     int
	to       int
	capacity int
	cost     int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const inf = 1 << 60

func minCostFlow(N int, edges []Edge, maxFlow int, src int, dst int) int {
	adj := make([][]int, N)
	cost := make([][]int, N)
	capacity := make([][]int, N)

	for i := range N {
		cost[i] = make([]int, N)
		capacity[i] = make([]int, N)
	}

	for _, e := range edges {
		adj[e.from] = append(adj[e.from], e.to)
		adj[e.to] = append(adj[e.to], e.from)
		cost[e.from][e.to] = e.cost
		cost[e.to][e.from] = -e.cost
		capacity[e.from][e.to] = e.capacity
	}

	d := make([]int, N)
	p := make([]int, N)
	marked := make([]bool, N)

	shortestPath := func() bool {
		for i := range N {
			d[i] = inf
			p[i] = -1
			marked[i] = false
		}
		var que []int
		que = append(que, src)
		marked[src] = true
		d[src] = 0

		for len(que) > 0 {
			u := que[0]
			que = que[1:]
			marked[u] = false
			for _, v := range adj[u] {
				if capacity[u][v] > 0 && d[v] > d[u]+cost[u][v] {
					d[v] = d[u] + cost[u][v]
					p[v] = u
					if !marked[v] {
						que = append(que, v)
						marked[v] = true
					}
				}
			}
		}
		return d[dst] != inf
	}

	var flow, res int
	for flow < maxFlow {
		if !shortestPath() {
			break
		}
		f := maxFlow - flow
		cur := dst
		for cur != src {
			f = min(f, capacity[p[cur]][cur])
			cur = p[cur]
		}

		flow += f
		res += f * d[dst]
		cur = dst

		for cur != src {
			capacity[p[cur]][cur] -= f
			capacity[cur][p[cur]] += f
			cur = p[cur]
		}
	}

	return res
}
