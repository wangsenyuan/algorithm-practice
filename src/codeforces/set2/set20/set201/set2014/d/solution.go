package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		ans := process(reader)
		buf.WriteString(fmt.Sprintf("%d %d\n", ans[0], ans[1]))
	}
	fmt.Print(buf.String())
}

func process(reader *bufio.Reader) []int {
	var n, d, k int
	fmt.Fscan(reader, &n, &d, &k)
	jobs := make([][]int, k)
	for i := range k {
		jobs[i] = make([]int, 2)
		fmt.Fscan(reader, &jobs[i][0], &jobs[i][1])
	}
	return solve(n, d, jobs)
}

func solve(n int, d int, jobs [][]int) []int {
	diff := make([]int, n+d+1)

	for _, job := range jobs {
		l, r := job[0], job[1]
		diff[l]++
		diff[r+d]--
	}
	ans := make([]int, 2)
	cnt := make([]int, 2)
	cnt[1] = n + 1

	var sum int
	for i := 1; i <= n; i++ {
		sum += diff[i]
		if i-d+1 > 0 {
			if sum > cnt[0] {
				cnt[0] = sum
				ans[0] = i - d + 1
			}
			if sum < cnt[1] {
				cnt[1] = sum
				ans[1] = i - d + 1
			}
		}
	}
	return ans
}

func solve1(n int, d int, jobs [][]int) []int {

	ans := make([]int, 2)
	cnt := make([]int, 2)
	cnt[1] = n + 1

	at := make([][]int, n)
	for _, job := range jobs {
		l, r := job[0], job[1]
		at[l-1] = append(at[l-1], r)
	}

	active := make(IntHeap, 0, n)

	for i := range n {
		for _, r := range at[i] {
			heap.Push(&active, r)
		}
		for len(active) > 0 && active[0] == i-d+1 {
			heap.Pop(&active)
		}
		if len(active) > cnt[0] {
			cnt[0] = len(active)
			ans[0] = max(0, i-d+1) + 1
		}
		if i >= d-1 && len(active) < cnt[1] {
			cnt[1] = len(active)
			ans[1] = i - d + 2
		}
	}

	return ans
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
