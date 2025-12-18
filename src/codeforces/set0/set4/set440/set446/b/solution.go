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
	var n, m, k, p int
	fmt.Fscan(reader, &n, &m, &k, &p)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(k, p, a)
}

func solve(k int, p int, a [][]int) int {
	n := len(a)
	m := len(a[0])
	row := make(IntHeap, n)
	col := make(IntHeap, m)
	for i := range n {
		for j := range m {
			row[i] += a[i][j]
			col[j] += a[i][j]
		}
	}
	heap.Init(&row)
	heap.Init(&col)

	dp := make([]int, k+1)

	for i := range k {
		dp[i+1] = dp[i] + row[0]
		v := heap.Pop(&row).(int)
		v -= m * p
		heap.Push(&row, v)
	}

	fp := make([]int, k+1)
	for i := range k {
		fp[i+1] = fp[i] + col[0]
		v := heap.Pop(&col).(int)
		v -= n * p
		heap.Push(&col, v)
	}

	best := -(1 << 60)

	for i := range k + 1 {
		best = max(best, dp[i]+fp[k-i]-i*(k-i)*p)
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
