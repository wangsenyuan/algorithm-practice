package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

const H = 11

func solve(k int, a []int) string {

	P := make([]int, H)
	P[0] = 1
	for i := 1; i < H; i++ {
		P[i] = P[i-1] * 10
	}

	a2 := make([][H]int, k)
	for i := range k {
		for j := range H {
			a2[i][j] = -1
		}
	}

	for _, v := range a {
		var l int
		for P[l] <= v {
			l++
		}
		a2[v%k][l] = v
	}

	dp := make([]*Item, k+1)
	pq := make(PriorityQueue, k+1)
	for i := range k + 1 {
		it := new(Item)
		it.id = i
		it.priority = inf
		it.index = i
		dp[i] = it
		pq[i] = it
	}
	dp[k].priority = 0
	heap.Init(&pq)

	type data struct {
		from int
		val  int
	}

	fa := make([]data, k+1)

	construct := func() string {
		var arr []int
		u := 0
		for u != k {
			arr = append(arr, fa[u].val)
			u = fa[u].from
		}
		slices.Reverse(arr)
		var buf strings.Builder
		for _, v := range arr {
			buf.WriteString(strconv.Itoa(v))
		}
		return buf.String()
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id
		if it.priority == inf {
			break
		}
		if u == 0 {
			return construct()
		}
		for i := range k {
			for j := range H {
				if a2[i][j] != -1 {
					v := (u*P[j]%k + a2[i][j]) % k
					if dp[v].priority > dp[u].priority+j {
						fa[v] = data{u, a2[i][j]}
						pq.update(dp[v], dp[u].priority+j)
					}
				}
			}
		}
	}
	return ""
}

const inf = 1 << 60

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
