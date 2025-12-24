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
	_, _, B := drive(reader)
	fmt.Println(B)
}

func drive(reader *bufio.Reader) (A int, c []int, B int) {
	var n int
	fmt.Fscan(reader, &n, &A)
	c = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}

	B = solve(A, c)
	return
}

func solve(A int, c []int) int {
	// n := len(c)

	x := max(A, slices.Max(c))

	freq := make([]int, x+1)

	items := make([]*Item, x+1)
	pq := make(PriorityQueue, x+1)
	for i := range x + 1 {
		it := new(Item)
		it.id = i
		it.priority = 0
		it.index = i
		pq[i] = it
		items[i] = it
	}

	items[A].priority = inf

	heap.Init(&pq)

	heap.Remove(&pq, items[0].index)

	for _, v := range c {
		freq[v]++
		if v != A && items[v].index >= 0 {
			pq.update(items[v], freq[v])
		}
		// 所有小于freq[A]的，都不能让B获胜
		for pq.Len() > 1 && pq[0].priority < freq[A] {
			heap.Pop(&pq)
		}
	}

	if pq.Len() > 1 {
		return pq[0].id
	}

	return -1
}

const inf = 1 << 60

// An Item is something we manage in a priority queue.
type Item struct {
	id       int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
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

func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
