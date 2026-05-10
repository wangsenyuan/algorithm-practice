package p3928

import (
	"container/heap"
)

const inf = 1 << 60

func minCost(n int, prices []int, roads [][]int) []int {
	adj := make([][]int, n)
	for i, e := range roads {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}

	nodes := make([]*node, n)
	for i := range n {
		nodes[i] = &node{i, inf, i}
	}

	dijkstra := func(start int) []int {
		pq := make(PriorityQueue, n)
		for i := range n {
			nodes[i].priority = inf
			pq[i] = nodes[i]
			pq[i].index = i
		}

		heap.Init(&pq)
		pq.update(nodes[start], 0)

		for pq.Len() > 0 {
			cur := heap.Pop(&pq).(*node)
			u := cur.id
			w := cur.priority
			for _, i := range adj[u] {
				v := roads[i][0] ^ roads[i][1] ^ u
				cost := roads[i][2]
				tax := roads[i][3]
				if nodes[v].priority > w+cost*tax {
					pq.update(nodes[v], w+cost*tax)
				}
			}
		}
		dist := make([]int, n)
		for i := range n {
			dist[i] = nodes[i].priority
		}
		return dist
	}

	dists := make([][]int, n)
	for i := range n {
		dists[i] = dijkstra(i)
	}

	play := func(s int) int {
		pq := make(PriorityQueue, n)
		for i := range n {
			nodes[i].priority = inf
			pq[i] = nodes[i]
			pq[i].index = i
		}
		heap.Init(&pq)
		pq.update(nodes[s], 0)
		for pq.Len() > 0 {
			cur := heap.Pop(&pq).(*node)
			u := cur.id
			w := cur.priority
			for _, i := range adj[u] {
				v := roads[i][0] ^ roads[i][1] ^ u
				nw := w + roads[i][2]
				if nw < nodes[v].priority {
					pq.update(nodes[v], nw)
				}
			}
		}

		ans := prices[s]

		for i := range n {
			ans = min(ans, prices[i]+nodes[i].priority+dists[i][s])
		}

		return ans
	}

	res := make([]int, n)
	for i := range n {
		res[i] = play(i)
	}
	return res
}

type node struct {
	id       int
	priority int
	index    int
}

type PriorityQueue []*node

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
	item := x.(*node)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *node, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
