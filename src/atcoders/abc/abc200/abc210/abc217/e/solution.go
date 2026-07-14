package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		var c int
		fmt.Fscan(reader, &c)
		if c == 1 {
			var x int
			fmt.Fscan(reader, &x)
			queries[i] = []int{c, x}
		} else {
			queries[i] = []int{c}
		}
	}
	return solve(queries)
}

func solve(queries [][]int) []int {
	var todo []int
	var pq IntHeap

	var res []int

	for _, cur := range queries {
		switch cur[0] {
		case 1:
			todo = append(todo, cur[1])
		case 2:
			if len(pq) > 0 {
				res = append(res, heap.Pop(&pq).(int))
			} else {
				res = append(res, todo[0])
				todo = todo[1:]
			}
		case 3:
			for len(todo) > 0 {
				heap.Push(&pq, todo[0])
				todo = todo[1:]
			}
		}
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
