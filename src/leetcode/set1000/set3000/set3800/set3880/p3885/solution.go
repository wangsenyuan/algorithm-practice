package p3885

import (
	"cmp"
	"container/heap"
)

type EventManager struct {
	pq     *PQ
	events map[int]*Event
}

func Constructor(events [][]int) EventManager {
	var pq PQ
	records := make(map[int]*Event)
	for _, evt := range events {
		cur := &Event{evt[0], evt[1], -1}
		heap.Push(&pq, cur)
		records[evt[0]] = cur
	}
	return EventManager{&pq, records}
}

func (this *EventManager) UpdatePriority(eventId int, newPriority int) {
	evt := this.events[eventId]
	evt.priority = newPriority
	heap.Fix(this.pq, evt.index)
}

func (this *EventManager) PollHighest() int {
	if this.pq.Len() == 0 {
		return -1
	}
	evt := heap.Pop(this.pq).(*Event)
	return evt.id
}

type Event struct {
	id       int
	priority int
	index    int
}

type PQ []*Event

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return cmp.Or(pq[j].priority-pq[i].priority, pq[i].id-pq[j].id) < 0
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	cur := x.(*Event)
	cur.index = len(*pq)
	*pq = append(*pq, cur)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	res.index = -1
	return res
}
