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
	free := k

	n := len(a)

	var id int
	var pq EventHeap
	running := make([]int, n)
	for i := range n {
		running[i] = -1
	}
	for free > 0 && id < n {
		heap.Push(&pq, Event{id: id, time: a[id]})
		running[id] = 0
		id++
		free--
	}

	marked := make([]bool, n)

	var m int
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(Event)
		running[cur.id] = -1
		m++
		free++
		// 展示的数据
		if id < n {
			heap.Push(&pq, Event{id: id, time: a[id] + cur.time})
			running[id] = cur.time
			id++
			free--
		}
		d := int(float64(100*m)/float64(n) + 0.5)
		for i := range id {
			// 在下一个任务完成前， 当前任务会展示出d
			if running[i] >= 0 && a[i] >= d && cur.time-running[i] < d &&
				(pq.Len() == 0 || pq[0].time >= running[i]+d) {
				marked[i] = true
			}
		}
	}

	var res int
	for i := range n {
		if marked[i] {
			res++
		}
	}
	return res
}

type Event struct {
	id   int
	time int
}

type EventHeap []Event

func (h EventHeap) Len() int { return len(h) }
func (h EventHeap) Less(i, j int) bool {
	return h[i].time < h[j].time ||
		h[i].time == h[j].time && h[i].id < h[j].id
}
func (h EventHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *EventHeap) Push(x any) {
	*h = append(*h, x.(Event))
}

func (h *EventHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
