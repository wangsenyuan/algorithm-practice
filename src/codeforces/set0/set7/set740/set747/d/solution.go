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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int {
	n := len(a)
	var neg []int
	for i := 0; i < n; i++ {
		if a[i] < 0 {
			neg = append(neg, i)
		}
	}
	if len(neg) == 0 {
		return 0
	}
	if len(neg) > k {
		return -1
	}
	res := 1
	m := len(neg)
	var sum int
	var pq IntHeap
	for i := 1; i < m; i++ {
		cur := neg[i] - neg[i-1] - 1
		sum += cur
		heap.Push(&pq, cur)

		for sum+i+1 > k {
			res += 2
			x := heap.Pop(&pq).(int)
			sum -= x
		}
	}

	sum += n - neg[m-1] - 1

	if sum+m > k {
		res++
	}

	return res
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
