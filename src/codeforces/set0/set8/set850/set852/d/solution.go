package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var V, E, N, K int
	fmt.Fscan(reader, &V, &E, &N, &K)

	start := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &start[i])
	}

	edges := make([][]int, E)
	for i := 0; i < E; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		edges[i] = []int{u, v, w}
	}

	return solve(V, start, edges, K)
}

const inf = 1 << 60

func solve(V int, start []int, edges [][]int, k int) int {
	g := NewGraph(V, len(edges)*2)
	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	items := make([]*Item, V)
	for i := range V {
		items[i] = &Item{
			id:       i,
			priority: 0,
			index:    i,
		}
	}

	dijkstra := func(s int) {
		pq := make(PriorityQueue, V)
		for i := range V {
			items[i].priority = inf
			items[i].index = i
			pq[i] = items[i]
		}

		items[s].priority = 0
		heap.Init(&pq)

		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			u := it.id

			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if items[v].priority > it.priority+g.val[i] {
					pq.update(items[v], it.priority+g.val[i])
				}
			}
		}
	}

	dist := make([][]int, V)
	for i := range V {
		dist[i] = make([]int, V)
		dijkstra(i)
		for j := range V {
			dist[i][j] = items[j].priority
		}
	}
	pair := make([]int, V)
	marked := make([]bool, V)
	n := len(start)
	adj := make([][]int, n + V)

	check := func(T int) bool {
		clear(adj)

		for i, s := range start {
			s--
			for j := range V {
				if dist[s][j] <= T {
					adj[i] = append(adj[i], j + n)
					adj[j+n] = append(adj[j+n], i)
				}
			}
		}

		for i := range V {
			pair[i] = -1
		}


		var dfs func(u int) bool

		dfs = func(u int) bool {
			// u < n
			for _, v := range adj[u] {
				if !marked[v - n] {
					marked[v - n] = true
					if pair[v - n] == -1 || dfs(pair[v-n]) {
						pair[v-n] = u
						return true
					}
				}
			}
			return false
		}

		var cnt int
		for i := range n {
			clear(marked)
			if dfs(i) {
				cnt++
				if cnt == k {
					return true
				}
			}
		}

		return false
	}

	res := sort.Search(1731312, check)
	if res == 1731312 {
		return -1
	}
	return res
}

type Item struct {
	id       int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[:n-1]
	it.index = -1
	return it
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
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
