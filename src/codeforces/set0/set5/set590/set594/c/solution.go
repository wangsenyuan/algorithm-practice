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
	magnents := make([][]int, n)
	for i := range n {
		magnents[i] = make([]int, 4)
		fmt.Fscan(reader, &magnents[i][0], &magnents[i][1], &magnents[i][2], &magnents[i][3])
	}
	return solve(k, magnents)
}

func solve(k int, magnents [][]int) int {
	n := len(magnents)
	var down MinPQ
	var up MaxPQ
	var left MinPQ
	var right MaxPQ

	items := make([][4]*Item, n)

	add := func(id int) {
		heap.Push(&left, items[id][0])
		heap.Push(&right, items[id][1])
		heap.Push(&down, items[id][2])
		heap.Push(&up, items[id][3])
	}

	for i := range n {
		x := magnents[i][0] + magnents[i][2]
		y := magnents[i][1] + magnents[i][3]

		items[i][0] = &Item{id: i, priority: x}
		items[i][1] = &Item{id: i, priority: x}
		items[i][2] = &Item{id: i, priority: y}
		items[i][3] = &Item{id: i, priority: y}

		add(i)
	}

	ans := inf

	update := func() {
		dx := right.Top().priority - left.Top().priority
		dy := up.Top().priority - down.Top().priority

		dx = max(1, (dx+1)/2)
		dy = max(1, (dy+1)/2)

		ans = min(ans, dx*dy)
	}

	remove := func(id int) {
		left.remove(items[id][0])
		right.remove(items[id][1])
		down.remove(items[id][2])
		up.remove(items[id][3])
	}

	var play0 func()
	var play1 func()
	var play2 func()
	var play3 func()

	play0 = func() {
		play1()
		var buf []int
		for left.Len()+k > n {
			id := left.Top().id
			buf = append(buf, id)
			remove(id)
			play1()
		}
		for _, id := range buf {
			add(id)
		}
	}

	play1 = func() {
		play2()
		var buf []int
		for right.Len()+k > n {
			id := right.Top().id
			buf = append(buf, id)
			remove(id)
			play2()
		}

		for _, id := range buf {
			add(id)
		}
	}

	play2 = func() {
		play3()
		var buf []int
		for down.Len()+k > n {
			id := down.Top().id
			buf = append(buf, id)
			remove(id)
			play3()
		}
		for _, id := range buf {
			add(id)
		}
	}

	play3 = func() {
		var buf []int
		for up.Len()+k > n {
			id := up.Top().id
			buf = append(buf, id)
			remove(id)
		}
		update()
		for _, id := range buf {
			add(id)
		}
	}

	play0()

	return ans
}

const inf = 1 << 60

type Item struct {
	id       int
	priority int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	old[n-1] = nil
	it.index = -1
	*pq = old[0 : n-1]
	return it
}

func (pq *PQ) Top() *Item {
	return (*pq)[0]
}

type MinPQ struct {
	PQ
}

func (pq *MinPQ) Less(i int, j int) bool {
	return pq.PQ[i].priority < pq.PQ[j].priority
}

func (pq *MinPQ) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *MinPQ) remove(it *Item) {
	oldPriority := it.priority
	pq.update(it, -inf)
	heap.Pop(pq)
	it.priority = oldPriority
}

type MaxPQ struct {
	PQ
}

func (pq *MaxPQ) Less(i int, j int) bool {
	return pq.PQ[i].priority > pq.PQ[j].priority
}

func (pq *MaxPQ) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *MaxPQ) remove(it *Item) {
	oldPriority := it.priority
	pq.update(it, inf)
	heap.Pop(pq)
	it.priority = oldPriority
}
