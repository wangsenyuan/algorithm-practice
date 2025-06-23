package p3594

import (
	"container/heap"
	"math/bits"
)

func minTime(n int, k int, m int, time []int, mul []float64) float64 {
	if n == 1 {
		return float64(time[0]) * mul[0]
	}
	if k == 1 {
		return -1
	}
	// n <= 12
	N := 1 << n

	dp := make([][]*Item, N)
	pq := make(PriorityQueue, N*m)
	for i := range N {
		dp[i] = make([]*Item, m)
		for j := range m {
			dp[i][j] = &Item{
				id:       i*m + j,
				priority: 1 << 60,
				index:    i*m + j,
			}
			pq[i*m+j] = dp[i][j]
		}
	}

	dp[0][0].priority = 0

	heap.Init(&pq)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		state := it.id / m
		stage := it.id % m

		if state == N-1 {
			return it.priority
		}

		left := N - 1 - state

		for sub := left; sub > 0; sub = (sub - 1) & left {
			d := bits.OnesCount(uint(sub))
			if d <= k {
				var cost int
				for i := range n {
					if (sub>>i)&1 == 1 {
						cost = max(cost, time[i])
					}
				}
				d := float64(cost) * mul[stage]
				ns := (stage + int(d)) % m
				next := state | sub

				if sub == left && dp[next][ns].priority > it.priority+d {
					dp[next][ns].priority = it.priority + d
					heap.Fix(&pq, dp[next][ns].index)
				}
				for j := range n {
					if (next>>j)&1 == 1 {
						d1 := float64(time[j]) * mul[ns]
						ns1 := (ns + int(d1)) % m
						tmp := next ^ (1 << j)
						if dp[tmp][ns1].priority > it.priority+d+d1 {
							dp[tmp][ns1].priority = it.priority + d + d1
							heap.Fix(&pq, dp[tmp][ns1].index)
						}
					}
				}
			}
		}
	}

	return -1
}

type Item struct {
	id       int
	priority float64
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
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
