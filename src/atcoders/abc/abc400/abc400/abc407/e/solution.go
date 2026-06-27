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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, 2*n)
	for i := range 2 * n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a []int) int {
	n := len(a)
	var res int
	var pq PQ
	// ()()()
	for i := range n {
		if i&1 == 0 {
			// 它现在是左括号, 但是尝试去替代一个前面的右括号, 它就可以作为右括号了
			if pq.Len() > 0 && pq[0].priority > a[i] {
				res += heap.Pop(&pq).(Item).priority
				heap.Push(&pq, Item{id: i, priority: a[i]})
			} else {
				res += a[i]
			}
		} else {
			heap.Push(&pq, Item{id: i, priority: a[i]})
		}
	}

	return res
}

type Item struct {
	id       int
	priority int
}

type PQ []Item

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority || (pq[i].priority == pq[j].priority && pq[i].id < pq[j].id)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func (pq *PQ) Pop() any {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}
