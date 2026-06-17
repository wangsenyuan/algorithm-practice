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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a []int, b []int) int {
	slices.Sort(a)
	slices.Sort(b)

	var res int
	var i int
	var pq IntHeap
	for _, v := range a {
		// 最好是正好吧 2 * v给使用掉
		for i < len(b) && b[i] <= 2*v {
			heap.Push(&pq, b[i])
			i++
		}
		if pq.Len() > 0 {
			res++
			heap.Pop(&pq)
		}
	}
	return res
}

type IntHeap []int

func (pq IntHeap) Len() int {
	return len(pq)
}

func (pq IntHeap) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq IntHeap) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *IntHeap) Push(x any) {
	*pq = append(*pq, x.(int))
}

func (pq *IntHeap) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}
