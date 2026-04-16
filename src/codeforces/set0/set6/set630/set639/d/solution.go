package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, k, b, c int
	fmt.Fscan(reader, &n, &k, &b, &c)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}
	return solve(k, t, b, c)
}

func solve(k int, t []int, b, c int) int {
	slices.Sort(t)

	if b >= 5*c {
		return solveCommentsOnly(k, t, c)
	}

	best := inf

	for y := range 5 {
		items := make([]pair, len(t))
		for i, v := range t {
			r := mod5(v)
			q := (v - r) / 5
			diff := (y - r + 5) % 5
			need := q
			if r > y {
				need++
			}
			weight := q*b - diff*c
			if r > y {
				weight += b
			}
			items[i] = pair{need, weight}
		}

		sort.Slice(items, func(i, j int) bool {
			return items[i].first < items[j].first
		})

		var pq IntHeap
		var sum int
		for _, it := range items {
			heap.Push(&pq, it.second)
			sum += it.second
			if pq.Len() > k {
				sum -= heap.Pop(&pq).(int)
			}
			if pq.Len() == k {
				best = min(best, k*it.first*b+sum*(-1))
			}
		}
	}

	return best
}

func solveCommentsOnly(k int, t []int, c int) int {
	var sum int
	best := inf
	for i, v := range t {
		if i >= k {
			sum -= t[i-k]
		}
		sum += v
		if i >= k-1 {
			best = min(best, (v*k-sum)*c)
		}
	}
	return best
}

func mod5(x int) int {
	x %= 5
	if x < 0 {
		x += 5
	}
	return x
}

type pair struct {
	first  int
	second int
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
	*h = old[:n-1]
	return x
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const inf = 1 << 60
