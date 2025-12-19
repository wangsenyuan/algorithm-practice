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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	voters := make([][]int, n)
	for i := range n {
		voters[i] = make([]int, 2)
		fmt.Fscan(reader, &voters[i][0], &voters[i][1])
	}
	return solve(voters)
}

func solve(people [][]int) int {
	slices.SortFunc(people, func(a, b []int) int {
		return b[1] - a[1]
	})

	var ids []int
	for _, cur := range people {
		ids = append(ids, cur[0])
	}

	slices.Sort(ids)
	ids = slices.Compact(ids)
	m := len(ids)
	arr := make([][]int, m)

	for _, cur := range people {
		j := sort.SearchInts(ids, cur[0])
		arr[j] = append(arr[j], cur[1])
	}

	check := func(k int) int {
		cnt := len(arr[0])
		var pq IntHeap
		var sum int
		for i := 1; i < m; i++ {
			for j := 0; j < min(k, len(arr[i])); j++ {
				heap.Push(&pq, arr[i][j])
			}
			for j := k; j < len(arr[i]); j++ {
				sum += arr[i][j]
				cnt++
			}
		}
		for cnt <= k && pq.Len() > 0 {
			sum += heap.Pop(&pq).(int)
			cnt++
		}
		return sum
	}

	l, r := 0, len(people)-len(arr[0])

	for l < r {
		mid := (l + r) >> 1
		x := check(mid)
		y := check(mid + 1)
		if x < y {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return check(l)
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
