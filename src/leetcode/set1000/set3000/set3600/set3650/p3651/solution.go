package p3651

import (
	"container/heap"
	"slices"
)

const inf = 1 << 60

func minCost(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	// 就运行k层就好了
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
		for j := range m {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0

	type data struct {
		r int
		c int
		v int
	}

	items := make([][]*Item, n)
	arr := make([]data, n*m)

	for i := range n {
		items[i] = make([]*Item, m)
		for j := range m {
			items[i][j] = &Item{i*m + j, inf, -1}
			arr[i*m+j] = data{i, j, grid[i][j]}
		}
	}

	slices.SortFunc(arr, func(a, b data) int {
		return a.v - b.v
	})

	for x := range k + 1 {
		pq := make(PriorityQueue, n*m)
		for i := range n {
			for j := range m {
				items[i][j].priority = dp[i][j]
				items[i][j].index = i*m + j
				pq[i*m+j] = items[i][j]
			}
		}

		heap.Init(&pq)

		var pos int

		for pq.Len() > 0 {
			cur := heap.Pop(&pq).(*Item)
			r, c := cur.id/m, cur.id%m
			if r+1 < n && items[r+1][c].priority > cur.priority+grid[r+1][c] {
				pq.update(items[r+1][c], cur.priority+grid[r+1][c])
				dp[r+1][c] = min(dp[r+1][c], cur.priority+grid[r+1][c])
			}
			if c+1 < m && items[r][c+1].priority > cur.priority+grid[r][c+1] {
				pq.update(items[r][c+1], cur.priority+grid[r][c+1])
				dp[r][c+1] = min(dp[r][c+1], cur.priority+grid[r][c+1])
			}
			for x < k && pos < n*m && arr[pos].v <= grid[r][c] {
				nr, nc := arr[pos].r, arr[pos].c
				dp[nr][nc] = min(dp[nr][nc], cur.priority)
				pos++
			}
		}
	}

	return dp[n-1][m-1]
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
