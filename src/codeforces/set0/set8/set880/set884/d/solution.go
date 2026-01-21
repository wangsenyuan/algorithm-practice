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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	if len(a) == 1 {
		return 0
	}
	if len(a) == 2 {
		return a[0] + a[1]
	}

	if len(a)%2 == 0 {
		a = append(a, 0)
	}

	pq := make(IntHeap, len(a))
	copy(pq, a)
	heap.Init(&pq)
	var res int
	for pq.Len() >= 3 {
		x := heap.Pop(&pq).(int)
		y := heap.Pop(&pq).(int)
		z := heap.Pop(&pq).(int)
		res += x + y + z
		heap.Push(&pq, x+y+z)
	}
	return res
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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
