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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Fprintf(writer, "No\n")
		return
	}
	fmt.Fprintf(writer, "Yes\n")
	for _, v := range res {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintf(writer, "\n")
}

func drive(reader *bufio.Reader) (islands [][]int, a []int, res []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	islands = make([][]int, n)
	for i := range n {
		islands[i] = make([]int, 2)
		fmt.Fscan(reader, &islands[i][0], &islands[i][1])
	}
	a = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a, islands)
	return

}

type bridge struct {
	id     int
	length int
}

type data struct {
	id   int
	dist int
}

func solve(a []int, islands [][]int) []int {
	n := len(islands)
	arr := make([]data, n-1)

	for i := 0; i < n-1; i++ {
		arr[i] = data{id: i, dist: islands[i+1][0] - islands[i][1]}
	}

	slices.SortFunc(arr, func(a, b data) int {
		return a.dist - b.dist
	})

	bridges := make([]bridge, len(a))
	for i, v := range a {
		bridges[i] = bridge{id: i, length: v}
	}

	slices.SortFunc(bridges, func(a, b bridge) int {
		return a.length - b.length
	})

	// 对于长度为x的桥来说，所有可以使用它的段中，应该选择那个最短的？
	res := make([]int, n-1)
	for i := range n - 1 {
		res[i] = -1
	}

	var long PriorityQueue

	m := len(a)

	for i, j := 0, 0; i < m; i++ {
		for j < n-1 && arr[j].dist <= bridges[i].length {
			u := arr[j].id
			heap.Push(&long, &Item{id: u, priority: islands[u+1][1] - islands[u][0]})
			j++
		}
		if long.Len() == 0 {
			continue
		}
		// bridges[i].length >= short[0].priority
		// 这里需要保证，那些能够使用的，都是  >= bridges[i].length 的部分
		// 但是这里存在一个问题，就是那个最大的被更早的使用掉了
		if long[0].priority < bridges[i].length {
			return nil
		}
		it := heap.Pop(&long).(*Item)
		j := it.id
		res[j] = bridges[i].id + 1
	}

	for i := 0; i < n-1; i++ {
		if res[i] < 0 {
			return nil
		}
	}
	return res
}

type Item struct {
	id       int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[:n-1]
	it.index = -1
	return it
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}

func (pq *PriorityQueue) remove(it *Item) {
	pq.update(it, -(1 << 60))
	heap.Pop(pq)
}
