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
	var n int
	fmt.Fscan(reader, &n)
	layers := make([][]int, n)
	for i := range n {
		var k, x1, x2 int
		fmt.Fscan(reader, &k, &x1, &x2)
		layers[i] = make([]int, k+3)
		layers[i][0] = k
		layers[i][1] = x1
		layers[i][2] = x2
		for j := range k {
			fmt.Fscan(reader, &layers[i][j+3])
		}
	}
	return solve(n, layers)
}

func solve(n int, layers [][]int) int {
	var pq1 PQ
	var pq2 LayerPQ

	lls := make([]*Layer, n)

	var x int
	for i, cur := range layers {
		k, x1, x2 := cur[0], cur[1], cur[2]
		x = max(x, x1)
		for j := range k {
			door := new(Door)
			door.layerId = i
			door.id = j
			door.left = x1
			door.width = cur[j+3]
			x1 += door.width
			heap.Push(&pq1, door)
		}
		// 这一层目前的最左端x1 - cur[1] 是所有门的宽度
		l := new(Layer)
		l.id = i
		l.left = x2 - (x1 - cur[1])
		lls[i] = l
		heap.Push(&pq2, l)
	}

	var best int

	for pq1.Len() > 0 {
		for pq1.Len() > 0 && pq1[0].left+pq1[0].width <= x {
			// 这个门可以移动到当前位置的后端
			it := heap.Pop(&pq1).(*Door)
			lid := it.layerId
			lls[lid].left += it.width
			heap.Fix(&pq2, lls[lid].index)
		}

		best = max(best, pq2[0].left-x)

		if pq1.Len() > 0 {
			// move to next position
			x = pq1[0].left + pq1[0].width
		}
	}

	return max(best, pq2[0].left-x)
}

type Door struct {
	layerId int
	id      int
	left    int
	width   int
	index   int
}

type PQ []*Door

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].left+pq[i].width < pq[j].left+pq[j].width
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PQ) Push(x any) {
	item := x.(*Door)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

type Layer struct {
	id    int
	left  int
	index int
}

type LayerPQ []*Layer

func (pq LayerPQ) Len() int {
	return len(pq)
}

func (pq LayerPQ) Less(i, j int) bool {
	return pq[i].left < pq[j].left
}

func (pq LayerPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *LayerPQ) Push(x any) {
	item := x.(*Layer)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *LayerPQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}
