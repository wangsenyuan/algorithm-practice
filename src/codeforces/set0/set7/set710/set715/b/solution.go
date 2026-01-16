package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, "YES")
	for _, e := range res {
		fmt.Fprintln(writer, e[0], e[1], e[2])
	}
}

func drive(reader *bufio.Reader) (n int, L int, s int, t int, res [][]int) {
	var m int
	fmt.Fscan(reader, &n, &m, &L, &s, &t)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	res = solve(n, edges, L, s, t)
	return
}

func solve(n int, edges [][]int, L int, s int, t int) [][]int {
	adj := make([][]int, n)

	var special []int
	m := len(edges)
	pos := make([]int, m)

	for i, e := range edges {
		adj[e[0]] = append(adj[e[0]], i)
		adj[e[1]] = append(adj[e[1]], i)
		if e[2] == 0 {
			special = append(special, i)
			pos[i] = len(special) - 1
		}
	}

	items := make([]*Item, n)

	for i := range n {
		it := new(Item)
		it.id = i
		it.priority = inf
		it.index = i
		items[i] = it
	}

	bfs := func(replace int, mid int) []int {
		pq := make(PriorityQueue, n)
		for i := range n {
			items[i].priority = inf
			items[i].index = i
			pq[i] = items[i]
		}
		items[s].priority = 0
		heap.Init(&pq)

		dist := make([]int, n)

		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			u := it.id
			dist[u] = it.priority
			for _, i := range adj[u] {
				v := edges[i][0] ^ edges[i][1] ^ u
				w := edges[i][2]
				if w == 0 {
					if pos[i] < mid {
						w = replace
					} else {
						w = inf
					}
				}
				if items[v].priority > it.priority+w {
					pq.update(items[v], it.priority+w)
				}
			}
		}
		return dist
	}

	d0 := bfs(inf, len(special))
	if d0[t] < L {
		return nil
	}
	d1 := bfs(1, len(special))
	if d1[t] > L {
		return nil
	}

	if len(special) == 0 {
		return edges
	}

	l, r := 0, len(special)

	for l < r {
		mid := (l + r) / 2
		if bfs(1, mid)[t] < L {
			r = mid
		} else {
			l = mid + 1
		}
	}
	r--
	for _, i := range special {
		if pos[i] < r {
			edges[i][2] = 1
		} else if pos[i] > r {
			edges[i][2] = inf
		}
	}

	d2 := bfs(1, len(edges))
	edges[special[r]][2] = L - d2[t] + 1

	return edges
}

const inf = 1000000000000000000

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
