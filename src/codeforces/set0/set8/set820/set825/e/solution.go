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
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) []int {
	adj := make([][]int, n)
	deg := make([]int, n)
	for _, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		adj[v] = append(adj[v], u)
		deg[u]++
	}

	var pq IntHeap

	for i := range n {
		if deg[i] == 0 {
			heap.Push(&pq, i)
		}
	}

	cur := n
	ans := make([]int, n)

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(int)
		ans[u] = cur
		cur--
		for _, v := range adj[u] {
			deg[v]--
			if deg[v] == 0 {
				heap.Push(&pq, v)
			}
		}
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
