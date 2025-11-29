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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int {
	n := len(a)

	// 舍弃最晚使用的书
	occ := make([]int, n+1)
	for i := range n + 1 {
		occ[i] = n
	}

	next := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		next[i] = occ[a[i]]
		occ[a[i]] = i
	}

	items := make([]*Item, n+1)

	for i := 1; i <= n; i++ {
		items[i] = new(Item)
		items[i].id = i
		items[i].priority = n
		items[i].index = -1
	}

	var res int
	var pq PriorityQueue

	for i := range n {
		v := a[i]
		if items[v].index >= 0 {
			// 书在手中
			pq.update(items[v], next[i])
			continue
		}
		if pq.Len() == k {
			// 必须舍弃一本书
			heap.Pop(&pq)
		}

		res++
		items[v].priority = next[i]
		heap.Push(&pq, items[v])

	}

	return res
}

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
	return pq[i].priority > pq[j].priority
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
func (pq *PriorityQueue) update(item *Item, p int) {
	item.priority = p
	heap.Fix(pq, item.index)
}
