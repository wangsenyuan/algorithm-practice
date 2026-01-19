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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, a, edges)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

type node struct {
	id  int
	val int
}

func solve(n int, a []int, edges [][]int) int {

	dp := make([]map[int]int, n)
	adj := make([][]node, n)

	items := make([]map[int]*Item, n)
	for i := range n {
		dp[i] = make(map[int]int)
		items[i] = make(map[int]*Item)
	}

	var pq PriorityQueue

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		dp[v][a[u]+a[v]] = add(dp[v][a[u]+a[v]], 1)
		adj[u] = append(adj[u], node{v, a[v]})
		if items[v][a[u]+a[v]] == nil {
			it := new(Item)
			it.id = v
			it.priority = a[u] + a[v]
			items[v][it.priority] = it
			heap.Push(&pq, it)
		}
	}

	for i := range n {
		slices.SortFunc(adj[i], func(a, b node) int {
			return a.val - b.val
		})
	}

	var res int

	for len(pq) > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id
		c := it.priority

		res = add(res, dp[u][c])

		for len(adj[u]) > 0 && adj[u][0].val < c {
			adj[u] = adj[u][1:]
		}

		for len(adj[u]) > 0 && adj[u][0].val == c {
			cur := adj[u][0]
			adj[u] = adj[u][1:]
			v := cur.id
			nc := a[u] + a[v]
			dp[v][nc] = add(dp[v][nc], dp[u][c])
			if items[v][nc] == nil {
				it := new(Item)
				it.id = v
				it.priority = nc
				items[v][nc] = it
				heap.Push(&pq, it)
			}
		}
	}

	return res
}

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
