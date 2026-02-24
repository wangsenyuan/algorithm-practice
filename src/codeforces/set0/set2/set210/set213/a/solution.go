package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	a := make([][]int, n)
	for i := range n {
		var k int
		fmt.Fscan(reader, &k)
		a[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(c, a)
}

func solve(c []int, a [][]int) int {
	n := len(c)
	adj := make([][]int, n)
	deg := make([]int, n)
	for i, v := range a {
		for _, u := range v {
			adj[u-1] = append(adj[u-1], i)
		}
		deg[i] = len(v)
	}

	items := make([]*Item, n)
	for i := range n {
		items[i] = new(Item)
		items[i].id = i
	}

	// 从s开始
	play := func(s int) int {
		pqs := make([]PriorityQueue, 3)
		for i := range n {
			it := items[i]
			it.priority = deg[i]
			heap.Push(&pqs[c[i]-1], it)
		}

		res := -1
		cnt := n
		for cnt > 0 {
			res++
			for pqs[s].Len() > 0 && pqs[s][0].priority == 0 {
				res++
				it := heap.Pop(&pqs[s]).(*Item)
				u := it.id
				cnt--
				for _, v := range adj[u] {
					pqs[c[v]-1].update(items[v], items[v].priority-1)
				}
			}
			s = (s + 1) % 3
		}
		return res
	}

	return min(play(0), play(1), play(2))
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
func (pq *PriorityQueue) update(item *Item, p int) {
	item.priority = p
	heap.Fix(pq, item.index)
}
