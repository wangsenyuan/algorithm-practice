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
	var n, x int
	fmt.Fscan(reader, &n, &x)
	candies := make([][]int, n)
	for i := range n {
		candies[i] = make([]int, 3)
		fmt.Fscan(reader, &candies[i][0], &candies[i][1], &candies[i][2])
	}
	return solve(x, candies)
}

type pair struct {
	first  int
	second int
}

func solve(x int, candies [][]int) int {
	arr := make([][]pair, 2)

	for _, cur := range candies {
		t := cur[0]
		// h, m
		arr[t] = append(arr[t], pair{cur[1], cur[2]})
	}

	// 按照高度递增
	for i := range 2 {
		slices.SortFunc(arr[i], func(a, b pair) int {
			return a.first - b.first
		})
	}

	play := func(a []pair, b []pair) int {
		// 先吃a，再吃b
		var pq1 IntHeap
		var pq2 IntHeap

		var res int

		cur := x

		for i, j := 0, 0; ; {
			for i < len(a) && a[i].first <= cur {
				heap.Push(&pq1, a[i].second)
				i++
			}
			if pq1.Len() == 0 {
				// can't process
				break
			}
			cur += heap.Pop(&pq1).(int)
			res++
			for j < len(b) && b[j].first <= cur {
				heap.Push(&pq2, b[j].second)
				j++
			}
			if pq2.Len() == 0 {
				break
			}
			cur += heap.Pop(&pq2).(int)
			res++
		}

		return res
	}

	ans1 := play(arr[0], arr[1])
	ans2 := play(arr[1], arr[0])

	return max(ans1, ans2)
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
