package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.10f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, m, r, k int
	fmt.Fscan(reader, &n, &m, &r, &k)
	return solve(n, m, r, k)
}

type pair struct {
	first  int
	second int
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(n int, m int, r int, k int) float64 {
	if n*m == k {
		// 每个格子里面都有鱼
		return float64(r) * float64(r)
	}

	n1 := (n - r + 1)
	m1 := (m - r + 1)

	getPriority := func(i int, j int) int {
		// 能够覆盖(i, j)的下网的数量
		t := max(0, i-r+1)
		l := max(0, j-r+1)
		i = min(i, n-r)
		j = min(j, m-r)
		return (i - t + 1) * (j - l + 1)
	}

	create := func(i int, j int) *Item {
		return &Item{
			id:       i*m + j,
			priority: getPriority(i, j),
		}
	}

	var pq PriorityQueue
	var sum int

	marked := make(map[pair]bool)
	marked[pair{n / 2, m / 2}] = true

	heap.Push(&pq, create(n/2, m/2))

	for k > 0 {
		k--
		it := heap.Pop(&pq).(*Item)
		// r, c 不是左上角
		r, c := it.id/m, it.id%m
		sz := it.priority
		// 它里面放一个鱼
		sum += sz

		for i := range 4 {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && !marked[pair{x, y}] {
				marked[pair{x, y}] = true
				heap.Push(&pq, create(x, y))
			}
		}
	}

	return float64(sum) / float64(n1*m1)
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
	return pq[i].priority > pq[j].priority
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
