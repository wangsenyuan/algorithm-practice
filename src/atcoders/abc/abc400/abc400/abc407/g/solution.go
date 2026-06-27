package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var h, w int
	fmt.Fscan(reader, &h, &w)
	a := make([][]int64, h)
	for i := range h {
		a[i] = make([]int64, w)
		for j := range w {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

func solve(a [][]int64) int64 {
	h, w := len(a), len(a[0])
	n := h * w
	s, t := n, n+1
	graph := NewMinCostFlow(n + 2)

	var sum int64
	for i := range h {
		for j := range w {
			sum += a[i][j]
			id := i*w + j
			if (i+j)&1 == 0 {
				graph.AddEdge(s, id, 1, 0)
			} else {
				graph.AddEdge(id, t, 1, 0)
			}
		}
	}

	dirs := []int{-1, 0, 1, 0, -1}
	for i := range h {
		for j := range w {
			if (i+j)&1 != 0 {
				continue
			}
			for k := 0; k < 4; k++ {
				x, y := i+dirs[k], j+dirs[k+1]
				if x < 0 || x == h || y < 0 || y == w {
					continue
				}
				cost := a[i][j] + a[x][y]
				if cost < 0 {
					graph.AddEdge(i*w+j, x*w+y, 1, cost)
				}
			}
		}
	}

	return sum - graph.MinCost()
}

const inf int64 = 1 << 60

type Edge struct {
	to, rev, cap int
	cost         int64
}

type MinCostFlow struct {
	g [][]Edge
}

func NewMinCostFlow(n int) *MinCostFlow {
	return &MinCostFlow{make([][]Edge, n)}
}

func (f *MinCostFlow) AddEdge(u, v, cap int, cost int64) {
	f.g[u] = append(f.g[u], Edge{v, len(f.g[v]), cap, cost})
	f.g[v] = append(f.g[v], Edge{u, len(f.g[u]) - 1, 0, -cost})
}

func (f *MinCostFlow) MinCost() int64 {
	n := len(f.g)
	dist := make([]int64, n)
	potential := make([]int64, n)
	prevv := make([]int, n)
	preve := make([]int, n)

	for i := range n {
		potential[i] = inf
	}
	potential[n-2] = 0
	queue := []int{n - 2}
	inQueue := make([]bool, n)
	inQueue[n-2] = true
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		inQueue[v] = false
		for _, e := range f.g[v] {
			if e.cap == 0 || potential[e.to] <= potential[v]+e.cost {
				continue
			}
			potential[e.to] = potential[v] + e.cost
			if !inQueue[e.to] {
				inQueue[e.to] = true
				queue = append(queue, e.to)
			}
		}
	}
	for i := range n {
		if potential[i] == inf {
			potential[i] = 0
		}
	}

	var res int64
	for {
		for i := range n {
			dist[i] = inf
		}
		dist[n-2] = 0
		pq := make(PriorityQueue, 1)
		pq[0] = Item{n - 2, 0}
		heap.Init(&pq)
		for pq.Len() > 0 {
			cur := heap.Pop(&pq).(Item)
			v := cur.node
			if dist[v] != cur.priority {
				continue
			}
			for i, e := range f.g[v] {
				if e.cap == 0 {
					continue
				}
				nd := dist[v] + e.cost + potential[v] - potential[e.to]
				if dist[e.to] <= nd {
					continue
				}
				dist[e.to] = nd
				prevv[e.to] = v
				preve[e.to] = i
				heap.Push(&pq, Item{e.to, nd})
			}
		}

		if dist[n-1] == inf || potential[n-1]+dist[n-1] >= 0 {
			break
		}
		for i := range n {
			if dist[i] < inf {
				potential[i] += dist[i]
			}
		}
		res += potential[n-1]
		for v := n - 1; v != n-2; v = prevv[v] {
			e := &f.g[prevv[v]][preve[v]]
			e.cap--
			f.g[v][e.rev].cap++
		}
	}

	return res
}

type Item struct {
	node     int
	priority int64
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
