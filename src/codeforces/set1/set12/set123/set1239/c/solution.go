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
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, p int
	fmt.Fscan(reader, &n, &p)
	t := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &t[i])
	}
	return solve(n, p, t)
}

func solve(n int, p int, t []int) []int {
	var events Events
	arr := make([]*Person, n)
	for i := range n {
		cur := Event{id: i, kind: 1, time: t[i]}
		heap.Push(&events, cur)
		arr[i] = &Person{id: i, index: -2}
	}

	var want People
	var que People
	var queueTime int

	ans := make([]int, n)

	playLeave := func(cur Event) {
		id := cur.id
		heap.Remove(&que, arr[id].index)
		ans[id] = cur.time
	}

	for len(events) > 0 {
		curTime := events[0].time
		for len(events) > 0 && events[0].time == curTime {
			cur := heap.Pop(&events).(Event)
			if cur.kind == 1 {
				heap.Push(&want, arr[cur.id])
			} else {
				// cur.kind == 0
				playLeave(cur)
			}
		}

		if len(want) > 0 {
			x := want[0].id
			if len(que) == 0 || que[0].id > x {
				heap.Pop(&want)
				nextTime := max(curTime, queueTime) + p
				event := Event{id: x, kind: 0, time: nextTime}
				heap.Push(&events, event)
				queueTime = nextTime
				heap.Push(&que, arr[x])
			}
		}
	}

	return ans
}

type Event struct {
	id    int
	kind  int
	time  int
	index int
}

type Events []Event

func (es Events) Len() int {
	return len(es)
}

func (es Events) Less(i, j int) bool {
	if es[i].time != es[j].time {
		return es[i].time < es[j].time
	}
	if es[i].kind != es[j].kind {
		return es[i].kind < es[j].kind
	}
	return es[i].id < es[j].id
}

func (es Events) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
	es[i].index = i
	es[j].index = j
}

func (es *Events) Push(x any) {
	e := x.(Event)
	e.index = len(*es)
	*es = append(*es, e)
}

func (es *Events) Pop() any {
	old := *es
	n := len(old)
	x := old[n-1]
	*es = old[0 : n-1]
	x.index = -1
	return x
}

type Person struct {
	id    int
	index int
}

type People []*Person

func (ps People) Len() int {
	return len(ps)
}

func (ps People) Less(i, j int) bool {
	return ps[i].id < ps[j].id
}

func (ps People) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
	ps[i].index = i
	ps[j].index = j
}

func (ps *People) Push(x any) {
	p := x.(*Person)
	p.index = len(*ps)
	*ps = append(*ps, p)
}

func (ps *People) Pop() any {
	old := *ps
	n := len(old)
	x := old[n-1]
	*ps = old[0 : n-1]
	x.index = -1
	return x
}
