package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var h, w, x int
	fmt.Fscan(reader, &h, &w, &x)
	var p, q int
	fmt.Fscan(reader, &p, &q)
	s := make([][]int64, h)
	for i := range h {
		s[i] = make([]int64, w)
		for j := range w {
			fmt.Fscan(reader, &s[i][j])
		}
	}
	return solve(x, p-1, q-1, s)
}

func solve(x, p, q int, s [][]int64) int64 {

	n := len(s)
	m := len(s[0])

	nodes := make([][]*node, n)

	marked := make([][]bool, n)
	for i := range n {
		nodes[i] = make([]*node, m)
		marked[i] = make([]bool, m)
		for j := range m {
			nodes[i][j] = &node{
				id:       i*m + j,
				index:    -1,
				priority: s[i][j],
			}
		}
	}
	nodes[p][q].priority = s[p][q]
	marked[p][q] = true

	var pq PriorityQueue
	heap.Push(&pq, nodes[p][q])

	var dd = []int{-1, 0, 1, 0, -1}

	var sum int64

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*node)
		r, c := cur.id/m, cur.id%m

		if sum > 0 && s[r][c] > (sum-1)/int64(x) {
			break
		}

		sum += s[r][c]

		for i := range 4 {
			nr, nc := r+dd[i], c+dd[i+1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && !marked[nr][nc] {
				marked[nr][nc] = true
				heap.Push(&pq, nodes[nr][nc])
			}
		}
	}

	return sum
}

type node struct {
	id       int
	priority int64
	index    int
}

type PriorityQueue []*node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
