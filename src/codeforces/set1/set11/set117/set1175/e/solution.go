package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	segs := make([][]int, n)
	for i := range n {
		segs[i] = make([]int, 2)
		fmt.Fscan(reader, &segs[i][0], &segs[i][1])
	}
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(segs, queries)
}

const FX = 500010

const INF = 1 << 60

func solve(segs [][]int, queries [][]int) []int {

	L := make([][]int, FX)
	R := make([][]int, FX)

	n := len(segs)
	items := make([]*Item, n)

	for i, cur := range segs {
		l, r := cur[0], cur[1]
		L[l] = append(L[l], i)
		R[r] = append(R[r], i)
		it := new(Item)
		it.id = i
		it.priority = r
		items[i] = it
	}

	var pq PriorityQueue

	h := 20

	fa := make([][]int, n)
	dp := make([][]int, n)
	for i := range n {
		fa[i] = make([]int, h)
		dp[i] = make([]int, h)
		dp[i][0] = segs[i][1]
	}

	a := make([][]int, FX)

	for i, cur := range queries {
		x := cur[0]
		a[x] = append(a[x], i)
	}

	m := len(queries)
	from := make([]int, m)

	for i := range m {
		from[i] = -1
	}

	adj := make([][]int, n)

	for i := range FX {
		for _, id := range L[i] {
			heap.Push(&pq, items[id])
		}

		if pq.Len() > 0 {
			for _, qid := range a[i] {
				from[qid] = pq[0].id
			}
		}

		for _, id := range R[i] {
			// 这个要离开了，需要找到最远的区间作为它的父节点，也可以是自己
			fa[id][0] = pq[0].id
			if fa[id][0] != id {
				adj[fa[id][0]] = append(adj[fa[id][0]], id)
			}
			heap.Remove(&pq, items[id].index)
		}
	}

	dep := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
			dp[u][i] = dp[fa[u][i-1]][i-1]
		}
		for _, v := range adj[u] {
			dep[v] = dep[u] + 1
			dfs(v)
		}
	}

	for i := range n {
		if fa[i][0] == i {
			dep[i] = 0
			dfs(i)
		}
	}

	find := func(qid int, x int, y int) int {
		// 要找到覆盖x的区间[l1, r1], 且r1最远
		u := from[qid]
		if u < 0 || dp[u][h-1] < y {
			return -1
		}

		v := u

		for i := h - 1; i >= 0; i-- {
			if dp[v][i] < y {
				v = fa[v][i]
			}
		}

		return dep[u] - dep[v] + 1
	}

	res := make([]int, len(queries))

	for i, cur := range queries {
		res[i] = find(i, cur[0], cur[1])
	}

	return res
}

type Item struct {
	id       int
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
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
