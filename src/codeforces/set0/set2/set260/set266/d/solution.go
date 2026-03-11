package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.3f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		for j := range 3 {
			fmt.Fscan(reader, &edges[i][j])
		}
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) float64 {
	type edge struct {
		to int
		w  int
	}

	adj := make([][]edge, n)
	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, w})
	}

	dijkstra := func(s int) []int {
		dist := make([]int, n)
		for i := range n {
			dist[i] = inf
		}
		dist[s] = 0
		pq := PQ{&Item{id: s, priority: 0}}
		heap.Init(&pq)

		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			u := it.id
			if it.priority != dist[u] {
				continue
			}
			for _, e := range adj[u] {
				nd := it.priority + e.w
				if nd < dist[e.to] {
					dist[e.to] = nd
					heap.Push(&pq, &Item{id: e.to, priority: nd})
				}
			}
		}

		return dist
	}

	dists := make([][]int, n)
	for i := range n {
		dists[i] = dijkstra(i)
	}

	check := func(limit int) bool {
		for u := range n {
			cur := 0
			for v := range n {
				cur = max(cur, 2*dists[u][v])
			}
			if cur <= limit {
				return true
			}
		}

		type interval struct{ l, r int }

		for _, e := range edges {
			u, v, w := e[0]-1, e[1]-1, 2*e[2]
			bad := make([]interval, 0, n)

			for i := range n {
				// Point t on the edge (0 <= t <= w, doubled):
				// min(2*du[i] + t, 2*dv[i] + w - t) <= limit
				// If both branches are > limit, t is forbidden.
				l := limit - 2*dists[u][i] + 1
				r := 2*dists[v][i] + w - limit - 1
				l = max(l, 0)
				r = min(r, w)
				if l <= r {
					bad = append(bad, interval{l, r})
				}
			}

			if len(bad) == 0 {
				return true
			}

			slices.SortFunc(bad, func(a, b interval) int {
				if a.l != b.l {
					return a.l - b.l
				}
				return a.r - b.r
			})

			covered := -1
			for _, seg := range bad {
				if seg.l > covered+1 {
					return true
				}
				covered = max(covered, seg.r)
				if covered >= w {
					break
				}
			}
			if covered < w {
				return true
			}
		}

		return false
	}

	hi := 0
	for u := range n {
		for v := range n {
			hi = max(hi, 2*dists[u][v])
		}
	}

	lo := 0
	for lo < hi {
		mid := (lo + hi) / 2
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return float64(lo) / 2
}

const inf = 1 << 60

type Item struct {
	id       int
	priority int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool { return pq[i].priority < pq[j].priority }

func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQ) Push(x any) { *pq = append(*pq, x.(*Item)) }

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
