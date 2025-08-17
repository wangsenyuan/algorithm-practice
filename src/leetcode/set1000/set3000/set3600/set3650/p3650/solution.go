package p3562

import "container/heap"

type pair struct {
	first  int
	second int
}

func minCost(n int, edges [][]int) int {
	adj := make([][]pair, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], pair{v, w})
		adj[v] = append(adj[v], pair{u, w * 2})
	}

	dist := make([]*Item, n)
	pq := make(PriorityQueue, n)
	for i := range n {
		dist[i] = &Item{i, inf, i}
		pq[i] = dist[i]
	}
	dist[0].priority = 0
	heap.Init(&pq)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id
		for _, e := range adj[u] {
			v, w := e.first, e.second
			if dist[v].priority > dist[u].priority+w {
				pq.update(dist[v], dist[u].priority+w)
			}
		}
	}

	if dist[n-1].priority == inf {
		return -1
	}
	return dist[n-1].priority
}

const inf = 1 << 60

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
