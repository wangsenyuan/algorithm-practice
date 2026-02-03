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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	songs := make([][]int, n)
	for i := range n {
		songs[i] = make([]int, 2)
		fmt.Fscan(reader, &songs[i][0], &songs[i][1])
	}
	return solve(k, songs)
}

const inf = 1 << 60

func solve(k int, songs [][]int) []int {
	var events Events

	n := len(songs)
	ans := make([]int, n)
	where := make([]int, n)

	for i := range n {
		where[i] = -1
		cur := Event{time: songs[i][0], id: i, kind: 1}
		heap.Push(&events, cur)
	}

	var servers IntHeap
	for range k {
		heap.Push(&servers, 0)
	}

	for events.Len() > 0 {
		cur := heap.Pop(&events).(Event)
		if cur.kind == 0 {
			ans[cur.id] = cur.time
		} else {
			// 应该找到最小的queTime
			first := heap.Pop(&servers).(int)
			next := Event{time: max(first, cur.time) + songs[cur.id][1], id: cur.id, kind: 0}
			heap.Push(&events, next)
			heap.Push(&servers, next.time)
		}
	}
	return ans
}

type Event struct {
	time int
	id   int
	kind int // 1 for arrive, 0 for leave
}

type Events []Event

func (this Events) Len() int {
	return len(this)
}

func (this Events) Less(i, j int) bool {
	return this[i].time < this[j].time || this[i].time == this[j].time && this[i].kind < this[j].kind
}

func (this Events) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *Events) Push(x any) {
	*this = append(*this, x.(Event))
}

func (this *Events) Pop() any {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0 : n-1]
	return x
}

type IntHeap []int

func (this IntHeap) Len() int {
	return len(this)
}

func (this IntHeap) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this IntHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *IntHeap) Push(x any) {
	*this = append(*this, x.(int))
}

func (this *IntHeap) Pop() any {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0 : n-1]
	return x
}
