package main

import (
	"bufio"
	"cmp"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, 2)
	for i := range 2 {
		a[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

func solve(a [][]int) int {
	n := len(a[0])
	pref := make([][]int, n)
	for i := range n {
		pref[i] = make([]int, 2)
		pref[i][0] = a[0][i]
		pref[i][1] = a[0][i]
		if i > 0 {
			pref[i][0] = min(a[0][i], pref[i-1][0])
			pref[i][1] = max(a[0][i], pref[i-1][1])
		}
	}

	suf := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		suf[i] = make([]int, 2)
		suf[i][0] = a[1][i]
		suf[i][1] = a[1][i]
		if i < n-1 {
			suf[i][0] = min(a[1][i], suf[i+1][0])
			suf[i][1] = max(a[1][i], suf[i+1][1])
		}
	}

	var arr [][]int

	for i := range n {
		// 如果在i处往下
		lo := min(pref[i][0], suf[i][0])
		hi := max(pref[i][1], suf[i][1])
		arr = append(arr, []int{lo, hi})
	}

	close := make([][]int, 2*n+1)
	for i, cur := range arr {
		hi := cur[1]
		close[hi] = append(close[hi], i)
	}
	var active IntHeap

	marked := make([]bool, len(arr))
	for i := range 2*n + 1 {
		for _, j := range close[i] {
			lo := arr[j][0]
			if len(active) > 0 && active[0] >= lo {
				continue
			}
			heap.Push(&active, lo)
			marked[j] = true
		}
	}
	var markedArr [][]int
	for i, cur := range arr {
		if marked[i] {
			markedArr = append(markedArr, cur)
		}
	}
	slices.SortFunc(markedArr, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})
	var res int
	for i := range len(markedArr) {
		l := markedArr[i][0]
		r := markedArr[i][1]
		nr := 2*n + 1
		if i+1 < len(markedArr) {
			nr = markedArr[i+1][1]
		}
		res += l * (nr - r)
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
