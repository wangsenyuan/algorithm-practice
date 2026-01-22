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
	var n, b int
	fmt.Fscan(reader, &n, &b)
	tasks := make([][]int, n)
	for i := range n {
		tasks[i] = make([]int, 2)
		fmt.Fscan(reader, &tasks[i][0], &tasks[i][1])
	}
	return solve(b, tasks)
}

// tasks is already sorted
func solve(b int, tasks [][]int) []int {
	var events Events
	for i, cur := range tasks {
		t := cur[0]
		event := Event{id: i, op: 1, time: t}
		heap.Push(&events, event)
	}

	n := len(tasks)
	ans := make([]int, n)

	var queSize int
	var queTime int

	b++

	for len(events) > 0 {
		cur := heap.Pop(&events).(Event)
		if cur.op == 1 {
			if queSize == b {
				ans[cur.id] = -1
				// reject it
				continue
			}
			// queSize < b
			queSize++
			// 等轮到这个task完成的时候的时间
			queTime = max(cur.time, queTime) + tasks[cur.id][1]
			event := Event{id: cur.id, op: 0, time: queTime}
			heap.Push(&events, event)
		} else {
			ans[cur.id] = cur.time
			queSize--
		}
	}
	return ans
}

type Event struct {
	id   int
	op   int
	time int
}

type Events []Event

func (events Events) Len() int {
	return len(events)
}

func (events Events) Less(i, j int) bool {
	return events[i].time < events[j].time || events[i].time == events[j].time && events[i].op < events[j].op
}

func (events Events) Swap(i, j int) {
	events[i], events[j] = events[j], events[i]
}

func (events *Events) Push(x any) {
	*events = append(*events, x.(Event))
}

func (events *Events) Pop() any {
	old := *events
	n := len(old)
	x := old[n-1]
	*events = old[0 : n-1]
	return x
}
