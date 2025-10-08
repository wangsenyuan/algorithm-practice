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
	ans := drive(reader)
	fmt.Fprintln(writer, ans)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	h := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &h[i])
	}
	return solve(n, h)
}

func solve(n int, h []int) int {
	// 还不能用bfs, 因为sum(h)太大了
	items := make([]*Item, n)
	pq := make(PriorityQueue, n)
	for i := range n {
		it := new(Item)
		it.id = i
		it.priority = h[i]
		if i == 0 || i == n-1 {
			// 两头的只需要一次就可以了
			it.priority = 1
		}
		it.index = i
		items[i] = it
		pq[i] = it
	}
	heap.Init(&pq)
	var res int
	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		res = it.priority
		i := it.id
		if i > 0 && items[i-1].index >= 0 && items[i-1].priority > it.priority+1 {
			pq.update(items[i-1], it.priority+1)
		}
		if i+1 < n && items[i+1].index >= 0 && items[i+1].priority > it.priority+1 {
			pq.update(items[i+1], it.priority+1)
		}
	}

	return res
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
