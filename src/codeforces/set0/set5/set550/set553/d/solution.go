package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, res := drive(reader)
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (n int, roads [][]int, F []int, res []int) {
	var m, k int
	fmt.Fscan(reader, &n, &m, &k)
	F = make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &F[i])
	}

	roads = make([][]int, m)
	for i := range m {
		roads[i] = make([]int, 2)
		fmt.Fscan(reader, &roads[i][0], &roads[i][1])
	}

	res = solve(n, roads, F)
	return
}

func solve(n int, roads [][]int, F []int) []int {

	fort := make([]bool, n)
	for _, v := range F {
		fort[v-1] = true
	}
	adj := make([][]int, n)
	for _, e := range roads {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	nodes := make([]*node, n)
	for i := range n {
		nodes[i] = &node{id: i, d1: len(adj[i]), d2: len(adj[i])}
	}

	play := func(d1 int, d2 int, pq *PQ) (int, int) {
		for i := range n {
			nodes[i].d1 = len(adj[i])
			if fort[i] {
				nodes[i].d1 = 0
			}
		}
		for _, u := range F {
			u--
			for _, v := range adj[u] {
				if !fort[v] {
					nodes[v].d1--
				}
			}
		}
		for i := range n {
			if !fort[i] {
				heap.Push(pq, nodes[i])
			}
		}

		c1, c2 := 0, 1

		for pq.Len() > 0 {
			it := (*pq)[0]
			if it.d1*c2 > c1*it.d2 {
				// 先更新最大值，再break
				c1, c2 = it.d1, it.d2
			}
			if it.d1*d2 >= it.d2*d1 {
				break
			}

			it = heap.Pop(pq).(*node)

			u := it.id
			for _, v := range adj[u] {
				if !fort[v] && nodes[v].index >= 0 {
					pq.update(nodes[v], nodes[v].d1-1, nodes[v].d2)
				}
			}

		}

		return c1, c2
	}

	d1, d2 := play(1, 1, &PQ{})

	var pq PQ

	play(d1, d2, &pq)

	var res []int
	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*node)
		res = append(res, it.id+1)
	}
	return res
}

type node struct {
	id    int
	d1    int
	d2    int
	index int
}

type PQ []*node

func (this PQ) Len() int {
	return len(this)
}

func (this PQ) Less(i, j int) bool {
	x := this[i].d1 * this[j].d2
	y := this[j].d1 * this[i].d2
	return x < y
}

func (this PQ) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *PQ) Push(x any) {
	it := x.(*node)
	it.index = len(*this)
	*this = append(*this, it)
}

func (this *PQ) Pop() any {
	old := *this
	n := len(old)
	it := old[n-1]
	it.index = -1
	*this = old[:n-1]
	return it
}

func (this *PQ) update(it *node, d1, d2 int) {
	it.d1 = d1
	it.d2 = d2
	heap.Fix(this, it.index)
}
