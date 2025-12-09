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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	envelopes := make([][]int, k)
	for i := 0; i < k; i++ {
		envelopes[i] = make([]int, 4)
		fmt.Fscan(reader, &envelopes[i][0], &envelopes[i][1], &envelopes[i][2], &envelopes[i][3])
	}
	return solve(n, m, envelopes)
}

const inf = 1 << 60

func solve1(n int, m int, envelopes [][]int) int {
	k := len(envelopes)

	items := make([]*Item, k)

	begin := make([][]int, n+1)
	end := make([][]int, n+1)

	for i := range k {
		it := new(Item)
		it.id = i
		s, t, d, w := envelopes[i][0], envelopes[i][1], envelopes[i][2], envelopes[i][3]
		it.w = w
		it.d = d
		items[i] = it
		begin[s] = append(begin[s], i)
		end[t] = append(end[t], i)
	}

	var pq PriorityQueue

	arr := make([]Item, n+1)

	for i := 1; i <= n; i++ {
		for _, j := range begin[i] {
			heap.Push(&pq, items[j])
		}
		if pq.Len() > 0 {
			arr[i] = *pq[0]
		} else {
			arr[i] = Item{d: i}
		}
		for _, j := range end[i] {
			pq.remove(items[j])
		}
	}

	dp := make([][]int, 2)
	for i := range 2 {
		dp[i] = make([]int, n+2)
		for j := range n + 2 {
			dp[i][j] = inf
		}
	}
	dp[0][1] = 0

	ans := inf

	for j := range m + 1 {
		for i := range n + 2 {
			dp[(j^1)&1][i] = inf
		}
		for i := 1; i <= n; i++ {
			dp[(j^1)&1][i+1] = min(dp[(j^1)&1][i+1], dp[j&1][i])
			dp[j&1][arr[i].d+1] = min(dp[j&1][arr[i].d+1], dp[j&1][i]+arr[i].w)
		}
		ans = min(ans, dp[j&1][n+1])
	}

	return ans
}

func solve(n int, m int, envelopes [][]int) int {
	dp := make([][]int, n+2)

	// k := len(envelopes)
	for i := range n + 2 {
		dp[i] = make([]int, m+1)
		for j := range m + 1 {
			dp[i][j] = inf
		}
	}

	k := len(envelopes)

	items := make([]*Item, k)

	begin := make([][]int, n+1)
	end := make([][]int, n+1)

	for i := range k {
		it := new(Item)
		it.id = i
		s, t, d, w := envelopes[i][0], envelopes[i][1], envelopes[i][2], envelopes[i][3]
		it.w = w
		it.d = d
		items[i] = it
		begin[s] = append(begin[s], i)
		end[t] = append(end[t], i)
	}

	dp[0][0] = 0

	var pq PriorityQueue

	for i := 0; i <= n; i++ {
		for _, j := range begin[i] {
			heap.Push(&pq, items[j])
		}

		if len(pq) > 0 {
			j := pq[0].id
			d, w := envelopes[j][2], envelopes[j][3]
			// 如果要干扰，就需要干扰这么久
			for j := 1; j <= m; j++ {
				dp[i+1][j] = min(dp[i+1][j], dp[i][j-1])
			}

			// 如果不干扰
			for j := range m + 1 {
				dp[d+1][j] = min(dp[d+1][j], dp[i][j]+w)
			}
		} else {
			for j := range m + 1 {
				dp[i+1][j] = min(dp[i+1][j], dp[i][j])
			}
		}
		for _, j := range end[i] {
			pq.remove(items[j])
		}
	}

	return slices.Min(dp[n+1])
}

// An Item is something we manage in a priority queue.
type Item struct {
	id    int // The value of the item; arbitrary.
	w     int
	d     int
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].w > pq[j].w || pq[i].w == pq[j].w && pq[i].d > pq[j].d
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) remove(it *Item) {
	w := it.w
	it.w = inf
	heap.Fix(pq, it.index)
	heap.Pop(pq)
	it.w = w
}
