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
	_, res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, cur := range res {
		fmt.Fprint(writer, len(cur))
		for _, v := range cur {
			fmt.Fprint(writer, " ", v)
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) (a [][]int, res [][]int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a = make([][]int, n)
	b := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
		b[i] = slices.Clone(a[i])
	}
	res = solve(b)
	return
}

func solve(a [][]int) [][]int {
	n := len(a)
	m := len(a[0])
	row := make([]int, n)
	col := make([]int, m)

	s1 := make([]*Item, n)
	s2 := make([]*Item, m)
	pq1 := make(PriorityQueue, n)
	pq2 := make(PriorityQueue, m)

	for i := range n {
		s1[i] = new(Item)
		s1[i].id = i
		s1[i].priority = 0
		s1[i].index = i
		pq1[i] = s1[i]
	}
	for i := range m {
		s2[i] = new(Item)
		s2[i].id = i
		s2[i].priority = 0
		s2[i].index = i
		pq2[i] = s2[i]
	}

	for i := range n {
		for j := range m {
			s1[i].priority += a[i][j]
			s2[j].priority += a[i][j]
		}
	}

	heap.Init(&pq1)
	heap.Init(&pq2)

	for pq1[0].priority < 0 || pq2[0].priority < 0 {
		if pq1[0].priority < 0 {
			r := pq1[0].id
			row[r] ^= 1
			for j := range m {
				pq1.update(s1[r], s1[r].priority-2*a[r][j])
				pq2.update(s2[j], s2[j].priority-2*a[r][j])
				a[r][j] = -a[r][j]
			}
			continue
		}
		c := pq2[0].id
		col[c] ^= 1
		for i := range n {
			pq1.update(s1[i], s1[i].priority-2*a[i][c])
			pq2.update(s2[c], s2[c].priority-2*a[i][c])
			a[i][c] = -a[i][c]
		}
	}

	res := make([][]int, 2)
	for i := range n {
		if row[i] == 1 {
			res[0] = append(res[0], i+1)
		}
	}
	for i := range m {
		if col[i] == 1 {
			res[1] = append(res[1], i+1)
		}
	}
	return res
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

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
