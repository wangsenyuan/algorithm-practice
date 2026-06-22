package p3970

import "container/heap"

func shortestPath(n int, edges [][]int, labels string, k int) int {
	adj := make([][]int, n)
	for i, e := range edges {
		u := e[0]
		adj[u] = append(adj[u], i)
	}

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, k+1)
		for j := range k + 1 {
			dist[i][j] = 1 << 60
		}
	}

	dist[0][1] = 0
	var pq PQ
	heap.Push(&pq, node{0, 1, 0})

	for len(pq) > 0 {
		cur := heap.Pop(&pq).(node)
		u, x := cur.u, cur.x
		if u == n-1 {
			return dist[u][x]
		}
		if dist[u][x] != cur.dist {
			continue
		}
		for _, eid := range adj[u] {
			v := u ^ edges[eid][0] ^ edges[eid][1]
			w := edges[eid][2]
			x1 := x
			if labels[u] == labels[v] {
				x1++
			} else {
				x1 = 1
			}
			if x1 <= k && dist[v][x1] > dist[u][x]+w {
				dist[v][x1] = dist[u][x] + w
				heap.Push(&pq, node{v, x1, dist[v][x1]})
			}
		}
	}

	return -1
}

type node struct {
	u    int
	x    int
	dist int
}

type PQ []node

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(node))
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	ret := old[n-1]
	*pq = old[:n-1]
	return ret
}
