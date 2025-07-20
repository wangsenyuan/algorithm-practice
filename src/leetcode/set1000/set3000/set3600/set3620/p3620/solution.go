package p3620

import (
	"container/heap"
	"sort"
)

const inf = 1 << 60

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)
	m := len(edges)
	g := NewGraph(n, m)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if online[u] && online[v] {
			g.AddEdge(u, v, w)
		}
	}

	items := make([]*Item, n)
	for i := range n {
		it := new(Item)
		it.id = i
		it.priority = inf
		items[i] = it
	}

	check := func(x int) bool {
		pq := make(PriorityQueue, n)
		for i := range n {
			it := items[i]
			it.priority = inf
			it.index = i
			pq[i] = it
		}
		items[0].priority = 0
		heap.Init(&pq)

		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			u := it.id
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				w := g.val[i]
				if w >= x && items[v].priority > items[u].priority+w {
					pq.update(items[v], items[u].priority+w)
				}
			}
		}

		return items[n-1].priority <= int(k)
	}
	if !check(0) {
		return -1
	}
	return sort.Search(1<<30, func(x int) bool {
		return !check(x)
	}) - 1
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
	next := make([]int, e+1)
	to := make([]int, e+1)
	val := make([]int, e+1)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
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
	old[n-1] = nil
	*pq = old[0 : n-1]
	return it
}

func (pq *PriorityQueue) update(it *Item, priority int) {
	it.priority = priority
	heap.Fix(pq, it.index)
}
