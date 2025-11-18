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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) string {
	var m, k int
	fmt.Fscan(reader, &m, &k)
	a := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &a[i])
	}
	passengers := make([][]int, m-1)
	for i := range m - 1 {
		passengers[i] = make([]int, 2)
		fmt.Fscan(reader, &passengers[i][0], &passengers[i][1])
	}
	return solve(a, passengers)
}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func solve(a []int, passengers [][]int) string {
	k := len(a)
	// m := len(passengers)

	items := make([]*Item, k)

	last_seen := make([]int, k)
	for i := range k {
		it := new(Item)
		it.id = i
		it.priority = a[i]
		items[i] = it
		last_seen[i] = -1
	}

	for i, cur := range passengers {
		if cur[0] > 0 {
			last_seen[cur[0]-1] = i
		}
	}

	var pq1 PriorityQueue

	buf := make([]byte, k)
	for i := range k {
		buf[i] = 'N'
	}

	for i := range k {
		if last_seen[i] == -1 {
			// 它就没有出现, 进入等待队列
			heap.Push(&pq1, items[i])
		}
	}

	var zero int
	var soldOut bool

	play := func() {
		var cnt int
		for pq1.Len() > 0 && pq1[0].priority <= zero {
			// 这些都可以用来消耗0
			it := heap.Pop(&pq1).(*Item)
			buf[it.id] = 'Y'
			// 将其移动到pq2中
			if !soldOut {
				soldOut = true
				cnt = it.priority
			}
		}

		zero -= cnt
	}

	for i, cur := range passengers {
		t, r := cur[0], cur[1]

		if r == 1 {
			// r == 1, 这里必须消耗一些
			play()
		}

		if t > 0 {
			t--
			a[t]--
			if last_seen[t] == i {
				// 后面不再有t了
				if a[t] == 0 {
					// 不用用 a[t] - cnt[t] <= zero 来判断， 因为它不一定能使用
					buf[t] = 'Y'
				}
				// 它可以用来消耗0
				items[t].priority = a[t]
				heap.Push(&pq1, items[t])
			}
		} else {
			zero++
		}
	}

	for i := range k {
		if a[i] <= zero {
			buf[i] = 'Y'
		}
	}

	return string(buf)
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
