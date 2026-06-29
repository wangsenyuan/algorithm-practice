package p3977

import "container/heap"

func minTimeMaxPower(n int, edges [][]int, power int, cost []int, source int, target int) []int64 {
	adj := make([][][]int, n)
	for _, e := range edges {
		u, v, t := e[0], e[1], e[2]
		adj[u] = append(adj[u], []int{v, t})
	}

	dp := make([][]int64, power+1)
	for i := range power + 1 {
		dp[i] = make([]int64, n)
		for j := range n {
			dp[i][j] = inf
		}
	}

	dp[power][source] = 0

	for x := power; x >= 0; x-- {
		for u, time := range dp[x] {
			if time == inf || x < cost[u] {
				continue
			}
			for _, to := range adj[u] {
				v, t := to[0], to[1]
				if dp[x-cost[u]][v] > time+int64(t) {
					dp[x-cost[u]][v] = time + int64(t)
				}
			}
		}
	}

	ans := []int64{inf, inf}

	for x := range power + 1 {
		if int64(dp[x][target]) <= ans[0] {
			ans[0] = int64(dp[x][target])
			ans[1] = int64(x)
		}
	}

	if ans[0] == inf {
		return []int64{-1, -1}
	}
	return ans
}

func minTimeMaxPower1(n int, edges [][]int, power int, cost []int, source int, target int) []int64 {
	adj := make([][][]int, n)
	for _, e := range edges {
		u, v, t := e[0], e[1], e[2]
		adj[u] = append(adj[u], []int{v, t})
	}
	// 现在要计算ans0, 就是要在最短时间能到达target
	// dp[u][x] 表示在到达位置u,且有x的能量时,花费的最少时间
	dp := make([][]*item, n)
	var pq PQ
	for i := range n {
		dp[i] = make([]*item, power+1)
		for j := range power + 1 {
			it := &item{
				id:       j*n + i,
				priority: inf,
			}
			heap.Push(&pq, it)
			dp[i][j] = it
		}
	}

	pq.update(dp[source][power], 0)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*item)
		u, x := it.id%n, it.id/n
		if u == target {
			break
		}
		// dp[u][x-1] 也需要被更新
		if x >= cost[u] {
			for _, to := range adj[u] {
				v, t := to[0], to[1]
				if dp[v][x-cost[u]].priority > it.priority+t {
					pq.update(dp[v][x-cost[u]], it.priority+t)
				}
			}
		}
	}

	ans := []int64{inf, inf}

	for x := range power + 1 {
		if int64(dp[target][x].priority) <= ans[0] {
			ans[0] = int64(dp[target][x].priority)
			ans[1] = int64(x)
		}
	}

	if ans[0] == inf {
		return []int64{-1, -1}
	}

	return ans
}

const inf = 1 << 60

type item struct {
	id       int
	priority int
	index    int
}

type PQ []*item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority || pq[i].priority == pq[j].priority && pq[i].id < pq[j].id
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	it := x.(*item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	res.index = -1
	*pq = old[:n-1]
	return res
}

func (pq *PQ) update(it *item, p int) {
	it.priority = p
	heap.Fix(pq, it.index)
}
