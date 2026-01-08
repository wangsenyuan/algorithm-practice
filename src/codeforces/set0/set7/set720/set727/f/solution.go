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

const inf = 1 << 60

func solve(a []int, b []int) []int {
	var pq IntHeap
	n := len(a)
	for i := n - 1; i >= 0; i-- {
		if a[i] < 0 {
			heap.Push(&pq, a[i])
		} else {
			for pq.Len() > 0 && a[i] >= 0 {
				a[i] += heap.Pop(&pq).(int)
			}
			if a[i] < 0 {
				heap.Push(&pq, a[i])
			}
		}
	}
	dp := make([]int, n+1)
	var cnt int
	for pq.Len() > 0 {
		cnt++
		dp[cnt] = heap.Pop(&pq).(int)
	}

	for i := 1; i <= cnt; i++ {
		dp[i] = -dp[i] + dp[i-1]
	}
	ans := make([]int, len(b))
	for i, v := range b {
		l, r := 1, cnt+1
		for l < r {
			mid := (l + r) / 2
			if dp[mid] > v {
				r = mid
			} else {
				l = mid + 1
			}
		}
		ans[i] = cnt - r + 1
	}

	return ans
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
