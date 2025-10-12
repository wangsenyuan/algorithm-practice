package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("No")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, "Yes")
	fmt.Fprintln(writer, len(res))
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) (a []int, res [][]int) {
	var n, s int
	fmt.Fscan(reader, &n, &s)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a)
	return
}

func solve(a []int) [][]int {
	var pq PriorityQueue
	for i, v := range a {
		it := new(Item)
		it.id = i + 1
		it.priority = v
		it.index = -1
		if v > 0 {
			heap.Push(&pq, it)
		}
	}

	var res [][]int

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		c := it.priority
		if c > len(pq) {
			return nil
		}
		var buf []*Item
		for range c {
			tmp := heap.Pop(&pq).(*Item)
			res = append(res, []int{it.id, tmp.id})
			tmp.priority--
			if tmp.priority > 0 {
				buf = append(buf, tmp)
			}
		}
		for _, tmp := range buf {
			heap.Push(&pq, tmp)
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
	return pq[i].priority > pq[j].priority
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
