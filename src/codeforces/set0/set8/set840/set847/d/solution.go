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
	var n, T int
	fmt.Fscan(reader, &n, &T)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}
	return solve(T, t)
}

type pair struct {
	first  int
	second int
}

func solve(T int, t []int) int {
	n := len(t)

	var best int

	var pq IntHeap
	for i := range n {
		delay := T - i - 2
		if delay < 0 {
			break
		}
		t[i] -= (i + 1)
		heap.Push(&pq, t[i])
		for pq.Len() > 0 && pq[0] > delay {
			heap.Pop(&pq)
		}
		best = max(best, pq.Len())
	}

	return best
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
