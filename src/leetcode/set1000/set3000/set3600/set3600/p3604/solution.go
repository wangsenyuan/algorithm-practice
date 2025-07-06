package p3604

import (
	"container/heap"
)

const inf = 1 << 60

func minTime(n int, edges [][]int) int {
	g := make([][]int, n)

	for i, e := range edges {
		u := e[0]
		g[u] = append(g[u], i)
	}

	items := make([]*Item, n)
	pq := make(PriorityQueue, n)
	for i := range n {
		it := new(Item)
		it.id = i
		it.priority = inf
		it.index = i
		items[i] = it
		pq[i] = it
	}

	items[0].priority = 0
	heap.Init(&pq)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id
		for _, eid := range g[u] {
			e := edges[eid]
			v, st, ed := e[1], e[2], e[3]
			if ed < it.priority {
				continue
			}
			next := max(st, it.priority) + 1
			if next < items[v].priority {
				pq.update(items[v], next)
			}
		}
	}

	res := items[n-1].priority
	if res == inf {
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
